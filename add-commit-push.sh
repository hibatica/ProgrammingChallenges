git add *

echo "Enter commit message:"

read COMMIT

git commit -m "$COMMIT"

git push

exit 0
