cmake_minimum_required(VERSION 3.14)
project(LeetCode_Cpp)

set(CMAKE_CXX_STANDARD 20)

include(FetchContent)
FetchContent_Declare(
  googletest
  URL https://github.com/google/googletest/archive/e2239ee6043f73722e7aa812a459f54a28552929.zip
)
# For Windows: Prevent overriding the parent project's compiler/linker settings
set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)
FetchContent_MakeAvailable(googletest)

enable_testing()

add_executable(1.TwoSum 0001.Two-Sum/1.TwoSum.cpp)
target_link_libraries(1.TwoSum gtest_main)

include(GoogleTest)
gtest_discover_tests(1.TwoSum)