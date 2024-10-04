#ifndef USER_H
#define USER_H
#include <string>
using namespace std;

class User{
private:
    string number;
    string name;
public:
    User(const string& name,const string& number);
    friend class LibrarySystem;
};


#endif // USER_H
