const fs = require("node:fs");
var os = require("os");

const processReports = (reports) => {
    let countSafeReport = 0;
    for (report of reports) {
        // check if report is safe
        const reportIsSafe = (report) => {
            // differences
            const levelDifferences = report
                .map((value, index) => report[index + 1] - value)
                .filter((value) => !isNaN(value));

            const allIncreasing = levelDifferences.every((value) => value > 0);
            const allDecreasing = levelDifferences.every((value) => value < 0);
            const allDifferences = levelDifferences
                .map((value) => Math.abs(value))
                .every((value) => value >= 1 && value <= 3);
            return (allIncreasing || allDecreasing) && allDifferences;
        };

        // process all reports
        isSafe = false;
        let n = 0;
        while (!isSafe && n < report.length) {
            // remove one
            const updatedReport = report.filter((_, index) => index !== n);
            isSafe = reportIsSafe(updatedReport);
            n++;
        }

        countSafeReport += isSafe;
    }

    console.log("Number of safe reports:", countSafeReport);
};

/*
const exampleData = [
    [7, 6, 4, 2, 1],
    [1, 2, 7, 8, 9],
    [9, 7, 6, 2, 1],
    [1, 3, 2, 4, 5],
    [8, 6, 4, 4, 1],
    [1, 3, 6, 7, 9],
];
processReports(exampleData);
*/

// load data
fs.readFile("data.txt", "utf8", (err, data) => {
    if (err) {
        console.error(err);
        return;
    }
    reports = data
        .split(os.EOL)
        .map((row) => row.split(" ").map((item) => Number(item)));

    processReports(reports);
});
