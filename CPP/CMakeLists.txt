cmake_minimum_required(VERSION 3.8)
set (CMAKE_CXX_STANDARD 14)
project(euler)

set (VERSION_MAJOR 1)
set (VERSION_MINOR 0)

set(BOOST_ROOT "/usr/local/boost-1.65.1")
set(Boost_USE_STATIC_LIBS        ON)  # only find static libs
set(Boost_USE_DEBUG_LIBS         OFF) # ignore debug libs and
set(Boost_USE_RELEASE_LIBS       ON)  # only find release libs
set(Boost_USE_MULTITHREADED      ON)
set(Boost_USE_STATIC_RUNTIME    OFF)

set(SOURCE main.cpp)

find_package(Boost 1.65.1 COMPONENTS thread regex system)
if(Boost_FOUND)
  include_directories(${Boost_INCLUDE_DIRS})
  add_executable(${PROJECT_NAME} ${SOURCE})
  target_link_libraries(${PROJECT_NAME} ${Boost_THREAD_LIBRARY} ${Boost_REGEX_LIBRARY} ${Boost_SYSTEM_LIBRARY})
endif()
