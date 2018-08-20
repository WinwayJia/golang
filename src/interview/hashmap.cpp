#include <stdio.h>
#include <stdint.h>
#include <iostream>
#include <unordered_map>

using namespace std;

struct St {
    string address;
    int age;
};

int main() {

    unordered_map<uint64_t, St> m = { {1, {"0000", 10 }, }};

    m[1].address = "11111";

    for (auto &i : m) {
        cout << i.second.address << "\t" << i.second.age << endl;
        printf("%p\n%p\n%p\n", &i, &(i.first), &(i.second));
    }

    return 0;
}
