#include <iostream>
#include <limits>
#include "useraccount.h"
#include "fixeddeposit.h"
#include "banksystem.h"

using namespace std;
string AccGen(const string& num){
    return "BSHDFC"+num;
}
bool validate(int& age){
    return age==0?false:true;
}

int main()
{
    BankSystem bankSystem;

    while (1)
    {
        cout << "Menu Options:" << endl;
        cout << "1. Add User Account" << endl;
        cout << "2. Open Fixed Deposit" << endl;
        cout << "3. Display User Accounts" << endl;
        cout << "4. Display Fixed Deposits" << endl;
        cout << "5. Exit" << endl;

        int choice;
        if (!(cin >> choice)) {
            cout << "Invalid input. Please enter a number between 1 and 5." << endl;
            cin.clear();
            cin.ignore(numeric_limits<streamsize>::max(), '\n');
            continue;
        }
        if (choice < 1 || choice > 5) {
            cout << "Wrong Input. Please enter a number between 1 and 5." << endl;
            continue;
        }

        switch (choice)
        {
        case 1:
        {
            string mobileNumber, name;
            int age;
            cout << "Enter Mobile Number: ";
            cin >> mobileNumber;
            cout << "Enter Name: ";
            cin >> name;
            cout << "Enter Age: ";
            cin >> age;
            if(!validate(age)){
                cout<<"Person opening account must be older than 1 year."<<endl;
            break;
            }
            UserAccount* user = new UserAccount(mobileNumber, name, age);
            bankSystem.addToBankRegister(user);
            break;
        }
        case 2:
        {
            string mobileNumber;
            cout<<"Enter Account only:"<<endl;
            cin >> mobileNumber;
            string accNum=AccGen(mobileNumber);
            if (bankSystem.searchAccount(accNum))
            {
                int amount, duration;
                cout << "Enter Account Number: ";
                cout << "Enter FD Amount: ";
                cin >> amount;
                cout << "Enter Duration (years): ";
                cin >> duration;
                UserAccount* user=bankSystem.searchUser(accNum);
                FixedDeposit* fd = new FixedDeposit(*user,amount, duration);
                bankSystem.addToFixedRegister(fd);
            }
            else
            {
                cout << "You Don't have an existing account!!\nPlease Provide details to Open FD:"<< endl;
                string name;
                int age;
                cout<<"------------------------------------------------------------"<<endl;
                cout<<"We took the above provided number as primary mobile number:"<<endl;
                cout << "Enter Name: ";
                cin >> name;
                cout << "Enter Age: ";
                cin >> age;
                if(!validate(age)){
                    cout<<"Person Opening account must be older than 1 year."<<endl;
                    break;
                }
                UserAccount* user = new UserAccount(mobileNumber, name, age);
                bankSystem.addToBankRegister(user);
                int amount, duration;
                cout << "Enter FD Amount: ";
                cin >> amount;
                cout << "Enter Duration (years): ";
                cin >> duration;
                FixedDeposit* fd = new FixedDeposit(*user,amount, duration);
                bankSystem.addToFixedRegister(fd);
            }
            break;
        }
        case 3:
            bankSystem.display();
            break;
        case 4:
            bankSystem.displayFD();
            break;
        case 5:
            cout << "Exiting..." << endl;
            exit(0);
        default:
            cout << "Invalid option, please try again." << endl;
        }

        cout << endl;
    }

    return 0;
}
