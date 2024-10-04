#ifndef BOOK_H
#define BOOK_H
#include <string>
using namespace std;


class Book{
private:
    string name;
    int quantity;
public:
    Book(string const& name,const int& quantity);
    friend class LibrarySystem;
};

#endif // BOOK_H
