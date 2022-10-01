COMMAND=gencode
COMMAND=genform
execute(){
	./${COMMAND} scenario constraint
	./${COMMAND} instance scope type definition campaign notes
	./${COMMAND} campaign version comments
	./${COMMAND} definition scenario scope type manualtcase1 manualtresult1 manualtcase2 manualtresult2 autotcase1 autotresult1 autotcase2 autotresult2
	./${COMMAND} testcase requirement approach method criteria
	./${COMMAND} requirement name description validation source
}
case "$1" in
e)	vi -p ${COMMAND}
	execute
	;;
"")	execute
	;;
esac
