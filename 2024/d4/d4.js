const fs = require("node:fs");
const os = require("os");

const word = "XMAS";

const findWords = (row) => {
    let rowString = "";
    row.forEach((char) => (rowString += char));
    let count = 0;
    let index = 0;
    while (index >= 0) {
        index = rowString.indexOf(word, index);
        if (index >= 0) {
            index = word.length + index;
            count++;
        }
    }
    return count;
};

const processLine = (line) => {
    // line
    let count = 0;
    count += findWords(line);

    // inverted
    const lineInverted = [];
    for (let j = line.length - 1; j >= 0; j--) {
        lineInverted.push(line[j]);
    }
    count += findWords(lineInverted);
    return count;
};

fs.readFile("input.txt", "utf-8", (error, data) => {
    const wordSearch = data.split(os.EOL).map((line) => line.split(""));

    const nRows = wordSearch.length;
    const nColumns = wordSearch[0].length;

    let totalCount = 0;

    // row by row
    for (let i = 0; i < nRows; i++) {
        let row = wordSearch[i];
        totalCount += processLine(row);
    }

    // column by column
    for (let j = 0; j < nColumns; j++) {
        let column = [];
        for (let i = 0; i < nRows; i++) {
            column.push(wordSearch[i][j]);
        }
        totalCount += processLine(column);
    }

    let i; //rows
    let j; //columns

    // diagonal: left side

    for (let n = 0; n < nRows; n++) {
        i = n;
        j = 0;
        let diagonal = [];
        while (i < nRows && j < nColumns && i >= 0 && j >= 0) {
            diagonal.push(wordSearch[i][j]);
            i++;
            j++;
        }
        totalCount += processLine(diagonal);
    }

    // diagonal: top side

    for (let m = 1; m < nColumns; m++) {
        j = m;
        i = 0;
        let diagonal = [];
        while (i < nRows && j < nColumns && i >= 0 && j >= 0) {
            diagonal.push(wordSearch[i][j]);
            i++;
            j++;
        }
        totalCount += processLine(diagonal);
    }

    // diagonal: right side

    for (let n = 0; n < nRows; n++) {
        i = n;
        j = nColumns - 1;
        let diagonal = [];
        while (i < nRows && j < nColumns && i >= 0 && j >= 0) {
            diagonal.push(wordSearch[i][j]);
            i--;
            j--;
        }
        totalCount += processLine(diagonal);
    }

    // diagnonal: bottom side

    for (let m = 0; m < nColumns - 1; m++) {
        j = m;
        i = nRows - 1;
        let diagonal = [];
        while (i < nRows && j < nColumns && i >= 0 && j >= 0) {
            diagonal.push(wordSearch[i][j]);
            i--;
            j--;
        }
        totalCount += processLine(diagonal);
    }

    console.log("Count", totalCount);
});
