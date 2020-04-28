/* Read input from STDIN. Print your output to STDOUT*/
#include <string>
#include <iostream>
#include <unordered_set>

using namespace std;

int main(int argc, char *a[])
{
    string line;
    int size;
    int vowelNumber;
    int consonantNumber;
    unordered_set<char> vowels ({'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'});
    unordered_set<char> consonants ({'b', 'B', 'c', 'C', 'd', 'D', 'f', 'F', 'g', 'G', 'h', 'H', 'j', 'J', 'k', 'K', 'l', 'L', 'm', 'M', 'n', 'N', 'p', 'P', 'q', 'Q', 'r', 'R', 's', 'S', 't', 'T', 'v', 'V', 'w', 'W', 'x', 'X', 'y', 'Y', 'z', 'Z'});

    getline(cin, line);
    try {
      size = stoi(line);
    } catch (std::invalid_argument const &e) {
      cout << "Bad input: std::invalid_argument thrown" << endl;
      exit(-1);
    } catch (std::out_of_range const &e) {
      cout << "Integer overflow: std::out_of_range thrown" << endl;
      exit(-1);
    }

    for (int i = 0; i < size; i++) {
        getline(cin, line);

        vowelNumber = 0;
        consonantNumber = 0;
        for (int j = 0; j < line.size(); j++) {
            auto searchVowels = vowels.find(line.c_str()[j]);
            if(searchVowels != vowels.end()) {
                vowelNumber++;
                continue;
            }
            auto searchConsonants = consonants.find(line.c_str()[j]);
            if (searchConsonants != consonants.end()) {
                consonantNumber++;
                continue;
            }
        }

        cout << vowelNumber << " " << consonantNumber << " " << vowelNumber * consonantNumber << endl;
    }
}
