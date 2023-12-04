t=require("fs").readFileSync("input.txt");o=14;
for(i=0;i<t.length;i++)if(new Set(t.subarray(i,i+o)).size==o){console.log(i+o);break}