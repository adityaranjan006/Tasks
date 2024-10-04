#include "banksystem.h"

BankSystem::BankSystem()
{

}
void BankSystem::addToBankRegister(UserAccount* user){
        bankRegister[user->accountNumber] = user;
    }

void BankSystem::addToFixedRegister(FixedDeposit* fd){
        fixedRegister[fd->getAccountNumber()].push_back(fd);
}

void BankSystem::display(){
    for (const auto& entry : bankRegister){
        cout << "Account Number: " << entry.first << ", Holder Name: " << entry.second->accountHolderName << endl;
    }
}

void  BankSystem::displayFD(){
        for (const auto& entry : fixedRegister){
                cout << "Account Number: " << entry.first << endl;
            for (const FixedDeposit* fd : entry.second){
                cout << "FD Amount: " << fd->getFAmount() << endl;
                cout << "FD Interest Rate: " << fd->getInterestRate() << "%" << endl;
                cout << "FD Maturity Amount:" << fd->FDAmount+(fd->FDAmount*fd->fdInterestRate*fd->duration)/100<<endl;
                cout<<endl;
            }
            cout << "----------------------------------" << endl;
        }
    }

UserAccount* BankSystem::searchUser(const string& accountNumber){
    auto it = bankRegister.find(accountNumber);
    if (it != bankRegister.end()){
        return it->second;
    }
        return nullptr;
}

bool BankSystem::searchAccount(const string& accNum){
    if(bankRegister.find(accNum)!=bankRegister.end()){
        return true;
    }
    return false;
}
bool BankSystem::mobileValidate(const string& mobileNumber){
    if(mobileNumber.length()!=10){
        return false;
    }
    return true;
}

void BankSystem::updateMobileNumber(const string& accNum){
    if(searchAccount(accNum)){
        string newNumber;
        cout<<"Enter the new Mobile Number:"<<endl;
        cin>>newNumber;
        if(mobileValidate(newNumber)){
            UserAccount* user=searchUser(accNum);
            user->accountNumber=newNumber;
            cout<<"Number updated for Account"<<accNum<<endl;
        }
        else{
            cout<<"Failed to Updated. Please Enter valid 10 digit mobile number"<<endl;
        }
    }
}
