assetNode start >~/.assetNode/log &
sleep 10
assetClient rest-server --chain-id test $1 $2>~/.assetClient/log &
echo "
Node and Client started up. For logs:
tail -f ~/.assetNode/log
tail -f ~/.assetClient/log
"
