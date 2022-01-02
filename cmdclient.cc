#include <iostream>
#include <chrono>
#include <thread>

static const int k_duration_seconds = 5 * 60; // 5 minutes

template <
    class result_t      = std::chrono::milliseconds,
    class clock_t       = std::chrono::steady_clock,
    class duration_t    = std::chrono::milliseconds
>
auto since(std::chrono::time_point<clock_t, duration_t> const& start) {
    return std::chrono::duration_cast<result_t>(clock_t::now() - start);
}

int main(int argc, char** argv)
{
    std::cout << "Take an Almont Time!\n\n";

    for (int i = 0; i < k_duration_seconds; ++i) {
        std::cout << i << " s\n";
        std::this_thread::sleep_for(std::chrono::milliseconds(1000));
    }

    std::cout << "\n\nGood job, well done!\n";

    return 0;
}
