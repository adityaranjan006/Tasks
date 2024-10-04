#ifndef LIBRARYSYSTEM_H
#define LIBRARYSYSTEM_H
#include <unordered_map>
#include <string>
#include <list>
#include "user.h"
#include "book.h"

using namespace std;

class LibrarySystem{
private:
    unordered_map<string,string> users;
    unordered_map<string,int> books;
    unordered_map<string,list<string>> borrowRegister;
public:
    void addUser(User& obj);
    void addBook(Book& obj);
    void viewUsers();
    void viewBooks();
    pair<bool,int> searchBook(string name);
    bool isValidUser(const string& number);
    void borrowBook(const string &number,const string& bookTitle);
    void returnBook(const string& userNumber, const string& bookTitle);
    void displayBorrowedBooks(const string& userNumber);
    bool passwordVerify(int pwd);
};

#endif // LIBRARYSYSTEM_H
