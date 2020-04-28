#include <iostream>
#include <random>
#include <climits>
#include <fstream>

using namespace std;

void generateTestCase(unsigned int teamSize, const unsigned long long int VALUE_UPPER_LIMIT);

int main(int argc, char** argv){

  // Errors and warning conditions
  if(argc == 1) {
    cout << "ERROR: argv[1] test case number!" << endl;
    exit(-1);
  }
  if(argc < 3) {
    cout << "WARNING: no team size was specified. It will be randomly choosen in [1, 10000]!" << endl;
  }
  if(argc < 4) {
    cout << "WARNING: No upper limit for values was specified. It will be set LLONG_MAX!" << endl;
  }
  
  // Parsing test case number
  unsigned int testCaseNumber = stoi(argv[1]);
  if(testCaseNumber <= 0 || testCaseNumber > 10000){
    cout << "Invalid test case size: " << argv[1] << endl;
    exit(-1);
  }

  // Getting team size
  unsigned int teamSize = (argc >= 3) ? stoi(argv[2]) : 0;

  // Getting value upper limit
  unsigned long long int valueUpperLimit = (argc == 4) ? stoi(argv[3]) :  (unsigned long long int)LLONG_MAX + 1; 
  
  // Redirecting cout
  ofstream out("agdata.txt");
  cout.rdbuf(out.rdbuf());

  // Generating test file
  cout << testCaseNumber << endl;
  for (int i = 0; i < testCaseNumber; i++) {
    generateTestCase(teamSize, valueUpperLimit);  
  }
}

void generateTestCase(unsigned int teamSize, const unsigned long long int VALUE_UPPER_LIMIT) {
  random_device rd;
  mt19937_64 eng(rd());
  uniform_int_distribution<unsigned long long int> distr;

  teamSize = (teamSize == 0) ? distr(eng) % 10000 : teamSize;
  cout << teamSize << endl;

  for(int i = 0; i < teamSize; i++){
    cout << distr(eng) % VALUE_UPPER_LIMIT << " ";
    if(i == teamSize - 1) {
      cout << endl;
    }
  }

  for(int i = 0; i < teamSize; i++){
    cout << distr(eng) % VALUE_UPPER_LIMIT << " ";
    if(i == teamSize - 1) {
      cout << endl;
    }
  }
}

