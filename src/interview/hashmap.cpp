#include <iostream>
#include <unordered_map>

using namespace std;

struct St {
    string address;
    int age;
};

int main() {

    unordered_map<string, St> m = { {"json", {"0000", 10 }, }};

    m["json"].address = "11111";

    for (auto i : m) {
        cout << i.second.address << "\t" << i.second.age << endl;
    }

    return 0;
}
