rm -rf Backend/web/admin/*
rm -rf Backend/web/front/*

cd web/frontend
yarn build
mv dist ../../Backend/web/front/

cd ../../web/admin
yarn build
mv dist ../../Backend/web/admin/

git add . 
git commit -m "update"
git push