/* Read input from STDIN. Print your output to STDOUT*/
#include <string>
#include <sstream>
#include <iostream>
#include <vector>
#include <climits>

using namespace std;

int getIngredientsNumber();
void parse(vector<unsigned long long int>& outputVector, string inputString, char delimiter);

int main(int argc, char *a[])
{
    unsigned int size, temp;
    unsigned long long int maxPpGirls = ULLONG_MAX;
    string quantityString;
    string availabilityString;
    vector<unsigned long long int> quantities;
    vector<unsigned long long int> availabilities;

    size = getIngredientsNumber();
    getline(cin, quantityString);
    getline(cin, availabilityString);
    parse(quantities, quantityString, ' ');
    parse(availabilities, availabilityString, ' ');

    for (int i = 0; i < size; i++) {
        temp = availabilities.at(i) / quantities.at(i);
        if(temp < maxPpGirls) {
            maxPpGirls = temp;
        }
    }

    cout << maxPpGirls << endl;
}

void parse(vector<unsigned long long int>& outputVector, string inputString, char delimiter) {
  stringstream ss(inputString); // Turn the string into a stream.
  string temp;
 
  while(getline(ss, temp, delimiter)) {
    outputVector.push_back(stoull(temp));
  }
}

int getIngredientsNumber() {
    string sizeString;
    unsigned int size = 0;

    try { 
      getline(cin, sizeString);
      size = stoi(sizeString);
    } catch (std::invalid_argument const &e) {
      cout << "Bad input: std::invalid_argument thrown" << endl;
      exit(-1);
    } catch (std::out_of_range const &e) {
      cout << "Integer overflow: std::out_of_range thrown" << endl;
      exit(-1);
    }
    return size;
}

