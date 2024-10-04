#ifndef BANKSYSTEM_H
#define BANKSYSTEM_H
#include <string>
#include <list>
#include <unordered_map>
#include "useraccount.h"
#include "fixeddeposit.h"

using namespace std;

class BankSystem
{
    private:
        unordered_map<string, UserAccount*> bankRegister;
        unordered_map<string,list<FixedDeposit*>> fixedRegister;
    public:
        BankSystem();
        UserAccount* searchUser(const string& accountNumber);
        void addToBankRegister(UserAccount* user);
        void addToFixedRegister(FixedDeposit* fd);
        void display();
        void displayFD();
        bool searchAccount(const string& accNum);
        void updateMobileNumber(const string& accNum);
        bool mobileValidate(const string& mobileNumber);
};

#endif // BANKSYSTEM_H
