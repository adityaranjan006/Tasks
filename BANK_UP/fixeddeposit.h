#ifndef FIXEDDEPOSIT_H
#define FIXEDDEPOSIT_H
#include <string>
#include <iostream>
#include "useraccount.h"
using namespace std;

class FixedDeposit
{
    private:
        string accountNumber;
        float fdInterestRate;
        int FDAmount;
        int duration;
        float interest;
    public:
        FixedDeposit(UserAccount& obj,int amount,int duration);
        string getAccountNumber() const;
        void openFD(UserAccount& obj,int amount,int duration);
        int getFAmount() const;
        float getInterestRate() const;
        friend class BankSystem;
};

#endif // FIXEDDEPOSIT_H
