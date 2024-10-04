#include "librarysystem.h"
#include <iostream>
using namespace std;

void LibrarySystem::addUser(User& obj)
{
    if(users.find(obj.number)!=users.end())
    {
        cout<<"Already Registered with Library!!"<<endl;
    }
    else
    {
        users[obj.number]=obj.name;
        cout<<"User "<<obj.name<< " Registered Successfully"<<endl;
    }
}
void LibrarySystem::addBook(Book& obj)
{
    if(books.find(obj.name)!=books.end())
    {
        //Already Registered
        cout<<"Already have the current Book"<<endl;
    }
    else
    {
        books[obj.name]=obj.quantity;
        cout<<"Book"<<obj.name<< "added Successfully to Inventory"<<endl;
    }
}
void LibrarySystem::viewUsers()
{
    for(auto it:users)
    {
        cout<<it.first<<" "<<it.second<<endl;
    }
}
void LibrarySystem::viewBooks()
{
    for(auto it:books)
    {
        cout<<it.first<<" "<<it.second<<endl;
    }
}

pair<bool,int> LibrarySystem::searchBook(string name)
{
    auto it=books.find(name);
    if(it!=books.end())
    {
        return make_pair(true,it->second);
    }
    return make_pair(false,0);
}
bool LibrarySystem::isValidUser(const string& number)
{
    if(users.find(number)!=users.end()) return true;
    return false;
}

void LibrarySystem::borrowBook(const string &number,const string& bookTitle)
{
    if(isValidUser(number))
    {
        pair<bool,int> res=searchBook(bookTitle);
        if(res.first && res.second>0)
        {
            borrowRegister[number].push_back(bookTitle);
            books[bookTitle]--;
            cout<<"Book "<<bookTitle<<" borrowed successfully"<<endl;
        }
        else if(res.second==0)
        {
            cout<<"Stock finished"<<endl;
        }
        else
        {
            cout<<"Invalid Book Title"<<endl;
        }
    }
    else
    {
        cout<<"Invalid User"<<endl;
    }
}

void LibrarySystem::returnBook(const string& userNumber, const string& bookTitle)
{
    auto it = borrowRegister.find(userNumber);
    if (it != borrowRegister.end())
    {
        auto& bks = it->second;
        bks.remove(bookTitle);
        books[bookTitle]++;
        cout << "Book" <<bookTitle<<" returned by user " << userNumber << endl;

    }
    else
    {
        cout << "User " <<userNumber<<" not found." << endl;
    }
}

void LibrarySystem::displayBorrowedBooks(const string& userNumber)
{
    auto it = borrowRegister.find(userNumber);
    if (it != borrowRegister.end())
    {
        cout << "Books borrowed by user " << userNumber << ":" << endl;
        for (auto& book : it->second)
        {
            cout << " - " << book << endl;
        }
    }
    else
    {
        cout << "User " << userNumber << " has not borrowed any books." << endl;
    }
}
