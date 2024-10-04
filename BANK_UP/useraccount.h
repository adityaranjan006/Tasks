#ifndef USERACCOUNT_H
#define USERACCOUNT_H
#include <string>
using namespace std;

class UserAccount
{
    private:
        string mobileNumber;
        string accountNumber;
        string accountType;
        string accountHolderName;
        string branch;
        int totalFDs;
        int age;
        int balance;
        float interestRate;
        bool gender;
    public:
        UserAccount(const string& mobileNumber,const string& name,int age);
        //Banksystem me bhi dal sakte h
        string getAccountNumber();
        string getAccountName();
        string getMobileNumber();
        friend class FixedDeposit;
        friend class BankSystem;
};

#endif // USERACCOUNT_H
