const fs = require("node:fs");

/* const memory =
    "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"; */

const re = /mul\((-?\d+(?:\.\d+)?),(-?\d+(?:\.\d+)?)\)/g;

const processMemory = (memory) => {
    const matches = [];
    let match;
    while ((match = re.exec(memory)) !== null) {
        matches.push([Number(match[1]), Number(match[2])]);
    }
    return matches.reduce((accumulator, currentValue) => {
        return (accumulator += currentValue[0] * currentValue[1]);
    }, 0);
};

/*
result = processMemory(memory);
console.log("result", result);
*/

fs.readFile("input.txt", "utf8", (err, data) => {
    if (err) {
        console.error(err);
        return;
    }
    const result = processMemory(data);
    console.log(result);
});
