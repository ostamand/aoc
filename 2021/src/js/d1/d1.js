const output = document.getElementById("output")
const select = document.querySelector("select.input")

document.getElementById("part-1").addEventListener("click", () => {
    getData()
        .then(data => {
            let nIncreases = 0
            for (let i = 1; i < data.length; i++) {
                if (data[i] > data[i - 1]) {
                    nIncreases++
                }
            }
            output.textContent = `Measurements larger: ${nIncreases}`
        })
})

document.getElementById("part-2").addEventListener("click", () => {
    getData()
        .then(data => {
            let sums = []
            for (let i = 0; i < data.length - 2; i++) {
                sums.push(data.slice(i, i + 3).reduce((total, current) => {
                    return total + current
                }, 0))
            }

            let nIncreases = 0
            for (let i = 1; i < sums.length; i++) {
                if (sums[i] > sums[i - 1]) {
                    nIncreases++
                }
            }

            output.textContent = `Sums larger: ${nIncreases}`
        })
})


function getData() {
    const filePath = `../../../inputs/${select.options[select.selectedIndex].text}`
    return fetch(filePath)
        .then(response => response.text())
        .then(text => processText(text))
}

/**
 * Get data from text
 * @param {string} text 
 */
function processText(text) {
    return text.split("\n").map(string => Number.parseInt(string))
}