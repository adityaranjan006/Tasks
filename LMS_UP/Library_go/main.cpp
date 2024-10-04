#include <iostream>
#include "librarysystem.h"
#include "user.h"
#include "book.h"
#include <cstdlib>
#include <unistd.h>

using namespace std;

int main()
{
    LibrarySystem lms;

    while(1)
    {
        int options;
        cout<<"1.Add Users"<<endl;
        cout<<"2.Add Books"<<endl;
        cout<<"3.View Books"<<endl;
        cout<<"4.View Users"<<endl;
        cout<<"5.Borrow Book"<<endl;
        cout<<"6.Return Book"<<endl;
        cout<<"7.View User Bucket"<<endl;
        cout<<"8.Exit Application"<<endl;
        cin>>options;
        switch(options)
        {
        case 1:
        {
            string name,number;
            cout<<"Enter Name of new User: ";
            cin>>name;
            cout<<"Enter Number of new User: ";
            cin>>number;
            User newUser(name,number);
            lms.addUser(newUser);
            sleep(1);
            system("cls");
        }
        break;
        case 2:
        {
            string name;
            int quantity;
            cout<<"Enter Name of new Book: ";
            cin>>name;
            cout<<"Enter quantity: ";
            cin>>quantity;
            Book newBook(name,quantity);
            lms.addBook(newBook);
            sleep(1);
            system("cls");
        }
        break;
        case 3:
        {
            lms.viewBooks();
        }
        break;
        case 4:
        {
            lms.viewUsers();
        }
        break;
        case 5:
        {
            string name,number;
            cout<<"Enter Name of Book: ";
            cin>>name;
            cout<<"Enter Number User: ";
            cin>>number;
            lms.borrowBook(number,name);
            sleep(1);
            system("cls");
        }
        break;
        case 6:
        {
            cout<<"For Returning a Book please provide:"<<endl;
            string name,number;
            cout<<"Enter Name of Book: ";
            cin>>name;
            cout<<"Enter Number User: ";
            cin>>number;
            lms.returnBook(number,name);
            sleep(1);
            system("cls");
        }
        break;
        case 7:
        {
            cout<<"See your Borrowed Books:"<<endl;
            string number;
            cout<<"Enter Number of User: ";
            cin>>number;
            lms.displayBorrowedBooks(number);
        }
        break;
        case 8:
        {
            exit(0);
        }
        break;
        default:
        {
            cout<<"Wrong Input"<<endl;
        }
        break;
        }
    }
    return 0;
}
