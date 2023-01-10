This repository is for testing out pull and merge event, rebase with master branch, change commit message, cherry pick from feature branch to master branch and drop commit

Commands Used:

git init
git add .
git commit -m "Initial Commit"
git remote add origin http://github.com/Abhishek-Mali-Simform/PracticalExam.git
git push -u origin master

//Branch Creation

git branch mybranch
git branch feature
git checkout mybranch
gedit mybranch.txt
git add .
git commit -m "mybranch commit"

// For pull Request

git push origin mybranch

//For Merge Branch

git checkout feature
gedit featureBranch.txt
git add .
git commit -m "featurebranch commit"
git checkout master
git pull //Pull used for previous push request(mybranch)
git merge feature
git push -u origin master

//Rebase 

gedit rebaseMaster.txt
git add .
git commit -m "add file form master"
git push origin master

git checkout mybranch
gedit rebasebranch.txt
git add .
git commit -m "add file from mybranch"
git push origin mybranch

git checkout master
git pull //Check for any updates
git rebase mybranch master
git push origin master // pushing rebase files

//change commit message

git checkout feature
git rebase -i HEAD~3
  // INSIDE log change 
  // pick 11c62ca featurebranch Commit to
  // reword 11c62ca Feature branch Commit
git push origin feature

//cherry pick
gedit README.md 
git add .
git commit -m "Add ReadMe.md File"
git push origin feature
git lof --oneline
git checkout master
git cherry-pick 5fcb4a
git push master origin

//Drop Commit
git reset --hard HEAD~1
//or
git rebase -i HEAD~3
  // INSIDE log change 
  // pick 11c62ca Feature branch Commit to
  // drop 11c62ca Feature branch Commit
 
 
