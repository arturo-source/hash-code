//Arturo Juan Martínez Sánchez 48791565D

#include <iostream>
#include <string.h>
#include <fstream>
#include <vector>
#include <algorithm>
#include <chrono>
#include <queue>
#include <tuple>

using namespace std;

//Elements of the problem
struct Problem {
    int maxValue;
    vector<int> values;
};

struct InfoNodes {
    int expandedNodes;
    int addedToAliveNodes;
    int notFeasibleNodes;
    int notPromisingNodes;
    int promisingNodesButDiscarded;
    int leafNodes;
    int currentBestValueUpdatedInLeaf;
    int currentBestValueUpdatedByPessimistic;
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
int greedy(Problem & problem, int k, int current) {
    for(unsigned int i=k; i<problem.values.size(); i++) {
        if(current + problem.values[i] <= problem.maxValue) {
            current += problem.values[i];
        }
    }

    return current;
}

int branchandbound(Problem & problem, InfoNodes & infoNodes) {
    typedef tuple<int, int, unsigned, int> node;
    struct is_worse {
        //sorted by optimistic bound
        bool operator() (const node& a, const node& b){
            return get<3>(a) < get<3>(b);
        }
    };
    priority_queue<node, vector<node>, is_worse> pq;
    int best_value = greedy(problem, 0, 0);
    int precalculatedRests = 0;
    
    for(unsigned int i=0; i<problem.values.size(); i++)
        precalculatedRests += problem.values[i];
    pq.push(node(0, precalculatedRests, 0, best_value));

    while(!pq.empty()){
        int value, rest, pessimistic;
        unsigned k;

        tie(value, rest, k, pessimistic) = pq.top();
        pq.pop();

        if(value+rest <= best_value) //not promising after best_value update 
            infoNodes.promisingNodesButDiscarded++;

        if(k==problem.values.size()) { //if it's a leaf
            if(value > best_value) {
                best_value = value;
                infoNodes.currentBestValueUpdatedInLeaf++;
            }
            infoNodes.leafNodes++;
            continue;
        }
        infoNodes.expandedNodes++;

        rest -= problem.values[k];
        for(unsigned i=0; i<2; i++) {
            int current_value = value + problem.values[k]*i;
            if(current_value==problem.maxValue) //the solution has been found
                return current_value;

            if(current_value <= problem.maxValue) { //if it's feasible
                pessimistic = greedy(problem, k+1, current_value);
                if(pessimistic > best_value) { //pessimistic is better
                    infoNodes.currentBestValueUpdatedByPessimistic++;
                    if(pessimistic==problem.maxValue)
                        return pessimistic;
                    best_value = pessimistic;
                }
                //it is only to sort the queue better
                pessimistic += current_value;
        
                int optimistic = current_value+rest;
                if(optimistic > best_value) { //if it's promising
                    pq.push(node(current_value, rest, k+1, pessimistic));
                    infoNodes.addedToAliveNodes++;
                } else { infoNodes.notPromisingNodes++; }
            } else { infoNodes.notFeasibleNodes++; }
        }
    }
    return best_value;
}

int main(int argc,char *argv[]) {
    bool errorReadingArguments = false;
    Problem problem{0, vector<int>()};
    InfoNodes infoNodes{0,0,0,0,0,0,0,0};

    checkArguments(problem, errorReadingArguments, argc, argv);

    if(errorReadingArguments) {
        cout << "Usage:" << endl << argv[0] << " -f file" << endl;
    } else {
        auto start = clock();
        //Sort the problem values to get a better result in greedy and faster solution in backtracking
        sort(problem.values.begin(), problem.values.end(), [&](int i, int j){return i>j;});
        int best_value = branchandbound(problem, infoNodes);
        auto end = clock();

        cout << "Best value: " << best_value << endl;
        cout << "CPU time (ms): " << 1000.0 * (end-start) / CLOCKS_PER_SEC << endl;
    }
}