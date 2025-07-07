const map = new Map();
let finished = false;

async function main() {
    const listPromises = []
    
    let MIN_ID = 0
    let MAX_ID = 100
    let STEP = 100

    while (true) {
        for (let i = MIN_ID; i < MAX_ID; i++) {
            
            listPromises.push(
                fetch(`http://localhost:8080/fragment?id=${i}`)
                    .then(r => r.json())
                    .then(j => saveText(j))
            )
        }
        await Promise.all(listPromises);

        MIN_ID += STEP;
        MAX_ID += STEP;
    }
}

function saveText(data) {
    map.set(data.index, data.text);
    validate();
}

function validate() {
    if (!finished && map.size >= 28) {
        finished = true;
        showResult();
        process.exit();
    }
}

function showResult() {
    const puzzle = 
        [...map.entries()]
        .sort(([a], [b]) => a - b)
        .map(([, value]) => value)
        .join(' ');

    console.log(puzzle)
}

main()