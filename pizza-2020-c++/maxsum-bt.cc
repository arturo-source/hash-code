#include <iostream>
#include <string.h>
#include <fstream>
#include <vector>
#include <algorithm>
#include <chrono>

using namespace std;

//Elements of the problem
struct Problem {
    int maxValue;
    vector<int> values;
};

//readFile returns true if there are any problems opening the file
bool readFile(string filePath, Problem & problem) {
    int valuesCount = 0;
    ifstream file;
    file.open(filePath, ios::out);
	
    if(file.is_open()) {
        file >> problem.maxValue;
        file >> valuesCount;

        for(int i=0;i<valuesCount; i++) {
            int value;
            file >> value;
            problem.values.push_back(value);
        }

        file.close();
        return false;
    } else {
        return true;
    }
}

//Turns true errorReadingArguments if there are errors.
void checkArguments(Problem & problem, bool & errorReadingArguments, int argc, char *argv[]) {
    if(argc == 1) {
        cout << "ERROR: missing arguments" << endl;
        errorReadingArguments = true;
    }
    for(int i=1; i<argc && !errorReadingArguments; i++) {
        if(!strcmp(argv[i],"-f")) {
            i++;
            if(i==argc) {
                cout << "ERROR: missing file path after -f argument" << endl;
                errorReadingArguments = true;
            } else {
                errorReadingArguments = readFile(argv[i], problem);
                if(errorReadingArguments) {
                    cout << "ERROR: can't open file: " << argv[i] << "." << endl;
                }
            }
        } else {
            cout << "ERROR: unknown option " << argv[i] << "." << endl;
            errorReadingArguments = true;
        }
    }
}

//returns the heuristic of take the most valued objects
int greedy(Problem & problem, vector<bool> &bestTakenObjects) {
    int solution = 0;
    for(unsigned int i=0; i<problem.values.size(); i++) {
        if(solution + problem.values[i] <= problem.maxValue) {
            solution += problem.values[i];
            bestTakenObjects[i] = true;
        }
    }

    return solution;
}

//backtracking calculates the best_value anyone can obtain with the restrictions
//and which objects are the chosen to this solution
void backtracking(Problem & problem, unsigned k, int current_value, int &best_value, 
        vector<bool> &takenObjects, vector<bool> &bestTakenObjects,
        int the_rest, bool &found) {
    //update solution (best_value and bestTakenObjects)
    if(k==problem.values.size() && current_value >= best_value) {
        best_value = current_value;
        bestTakenObjects = takenObjects;
        return;
    }
    for(unsigned i=0; i<2; i++) {
        takenObjects[k] = i;
        int optimistic = current_value + the_rest;
        current_value += problem.values[k] * i;
        //if is feasible and promising
        if (current_value <= problem.maxValue && optimistic > best_value) {
            found = current_value == problem.maxValue;
            //update solution (best_value and bestTakenObjects)
            if(current_value > best_value || found) {
                best_value = current_value;
                bestTakenObjects = takenObjects;
            }
            backtracking(problem, k+1, current_value, best_value, takenObjects, bestTakenObjects, the_rest-problem.values[k], found);
        }
        if(found) break;
    }
}

int main(int argc,char *argv[]) {
    bool errorReadingArguments = false;
    Problem problem{0, vector<int>()};

    checkArguments(problem, errorReadingArguments, argc, argv);

    if(errorReadingArguments) {
        cout << "Usage:" << endl << argv[0] << " -f file" << endl;
    } else {
        auto start = clock();
        int precalculatedRests = 0;
        for(unsigned int i=0; i<problem.values.size(); i++)
            precalculatedRests += problem.values[i];

        bool found = false;
        vector<bool> takenObjects(problem.values.size());
        vector<bool> bestTakenObjects(problem.values.size());
        
        //Sort the problem values to get a better result in greedy and faster solution in backtracking
        sort(problem.values.begin(), problem.values.end(), [&](int i, int j){return i>j;});
        int best_value = greedy(problem, bestTakenObjects);
        backtracking(problem, 0, 0, best_value, takenObjects, bestTakenObjects, precalculatedRests, found);
        auto end = clock();

        cout << "Best value: " << best_value << endl;
        cout << "Selection: ";
        for(unsigned int i=0; i<bestTakenObjects.size(); i++)
            if(bestTakenObjects[i])
                cout << problem.values[i] << " ";
        cout << endl;
        cout << "CPU time (ms): " << 1000.0 * (end-start) / CLOCKS_PER_SEC << endl;
    }
}