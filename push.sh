git add .
read -r -p "Your email: " email
read -r -p "Your name: " mail
git config --global user.email "$email"
git config --global user.name "$mail"
read -r -p "Your commit message: " message
git commit -m "$message"
git push