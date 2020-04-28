/* Read input from STDIN. Print your output to STDOUT*/
#include <string>
#include <sstream>
#include <iostream>
#include <algorithm>
#include <vector>
#include <climits>

using namespace std;

void swap(vector<unsigned long long int>& gRevolutionTeam, 
          unsigned int i, 
          unsigned int j);

void computeIthElement(vector<unsigned long long int>& gRevolutionTeam, 
                       unsigned int targetIndex, 
                       unsigned long long int enemyValue);

unsigned int getWins(vector<unsigned long long int>& gRevolutionTeam, 
                     vector<unsigned long long int>& enemyTeam, 
                     unsigned int size);

void getTeamPower(vector<unsigned long long int>& outputVector, 
                  string inputString, 
                  char delimiter);

int getSize();

int main(int argc, char *a[])
{
    unsigned int matchNumber, teamSize, maxWins, tempWins;
    string gRevolutionString;
    string enemyString;
    vector<unsigned long long int> gRevolutionTeam;
    vector<unsigned long long int> enemyTeam;

    matchNumber = getSize();

    for (int i = 0; i < matchNumber; i++) {
        teamSize = getSize();
        getline(cin, gRevolutionString);
        getline(cin, enemyString);
        getTeamPower(gRevolutionTeam, gRevolutionString, ' ');
        getTeamPower(enemyTeam, enemyString, ' ');

        for(int j = 0; j < gRevolutionTeam.size(); j++) {
            computeIthElement(gRevolutionTeam, j, enemyTeam.at(j));
        }

        cout << getWins(gRevolutionTeam, enemyTeam, teamSize) << endl;
    }
}

void swap(vector<unsigned long long int>& gRevolutionTeam, unsigned int i, unsigned int j) {
    if (i == j) {
        return;
    }
    
    unsigned long long int temp = gRevolutionTeam.at(i);
    gRevolutionTeam.at(i) = gRevolutionTeam.at(j);
    gRevolutionTeam.at(j) = temp;
}

void computeIthElement(vector<unsigned long long int>& gRevolutionTeam, 
                               unsigned int targetIndex, 
                               unsigned long long int enemyValue) {
    bool isWon = false;
    unsigned int winnerIndex, looserIndex;
    unsigned long long int gRevolutionValue;
    unsigned long long int delta = ULLONG_MAX;
    unsigned long long int minValue = ULLONG_MAX;

    for (int i =  targetIndex; i < gRevolutionTeam.size(); i++) {
        gRevolutionValue = gRevolutionTeam.at(i);

        if(gRevolutionValue > enemyValue) {
            if( (gRevolutionValue - enemyValue) < delta) {
                winnerIndex = i;
                delta = gRevolutionValue - enemyValue;
                isWon = true;
            }
        }

        if(gRevolutionValue < minValue) {
            minValue = gRevolutionValue;
            looserIndex = i;
        }
    }

    if(isWon) {
        swap(gRevolutionTeam, targetIndex, winnerIndex);
    } else {
        swap(gRevolutionTeam, targetIndex, looserIndex);
    }
}

unsigned int getWins(vector<unsigned long long int>& gRevolutionTeam, vector<unsigned long long int>& enemyTeam, unsigned int size) {
    unsigned int wins = 0;
    
    for (int i = 0; i < size; i++) {
        if (gRevolutionTeam.at(i) > enemyTeam.at(i)) {
            wins++;
        } 
    }
    return wins;
}

void getTeamPower(vector<unsigned long long int>& outputVector, string inputString, char delimiter) {
  stringstream ss(inputString); // Turn the string into a stream.
  string temp;
 
  outputVector.clear();
  while(getline(ss, temp, delimiter)) {
    outputVector.push_back(stoull(temp));
  }
}

int getSize() {
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
