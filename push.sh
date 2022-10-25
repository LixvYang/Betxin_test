git add .
git commit -m "update"
git push


git clone https://github.com/LixvYang/Betxin.git BetxinCopy
rm -rf Betxin/
mv BetxinCopy/ Betxin/
cp -r config/ Betxin/Backend/
mkdir /Betxin/Backend/log
cd Betxin/Backend/

echo "请关闭8080端口 lsof -i:8080"
echo "并再次启用8080"
echo "go run main.go &"
~                         