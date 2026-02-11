# Encode And Decode Stream

Same encoding rules as **Encode And Decode**, but the input arrives as a stream of values one at a time instead of a complete array.

## Encoding Rules (recap)

- **RLE Run:** Consecutive identical values with count >= 8 → `"RLE[value,count]"`.
- **BP Run:** Groups of up to 8 non-RLE values → `"BP[v1,v2,...,vk]"`. Exactly 8 per run except the last may have fewer.
- **Last run exception:** The final run may be emitted as RLE even if its count < 8, provided the BP buffer is empty.

## Streaming API

### Encoder

- `Write(value) → []string` — feeds one integer into the encoder. Returns any encoded runs that are now complete (may be empty).
- `Flush() → []string` — signals end of stream. Returns any remaining encoded runs.

### Decoder

- `Write(run) → []int` — feeds one encoded run string. Returns the decoded integers from that run.

The concatenation of all outputs from `Write` + `Flush` must equal the batch `Encode` result, and the decoded values must equal the original stream.

## Example

```
enc = new StreamEncoder()
enc.Write(5) → []          // building run of 5s
enc.Write(5) → []
enc.Write(5) → []
enc.Write(5) → []
enc.Write(5) → []
enc.Write(5) → []
enc.Write(5) → []
enc.Write(5) → []          // run of 5s reaches 8, but not finalized yet
enc.Write(1) → ["RLE[5,8]"] // new value → previous run (count=8) emitted as RLE
enc.Write(2) → []
enc.Write(3) → []
enc.Flush()  → ["BP[1,2,3]"]  // remaining BP buffer flushed

dec = new StreamDecoder()
dec.Write("RLE[5,8]") → [5,5,5,5,5,5,5,5]
dec.Write("BP[1,2,3]") → [1,2,3]
```
