window.onload = function() {

    //var mydata = JSON.parse(dictionary);
    console.log(dictionary);

    var chartData = {}

    var rawLables = Object.keys(dictionary)
    var rawDataset = []
    var rawDatasetDictionary = {}
    var rawData = Object.values(dictionary)

    var backgroundColor = []
    for (var i = 0; i < rawLables.length; i++) {
        newColor = getRandomColor();
        backgroundColor.push(newColor)
    }

    // Datasets
    rawDatasetDictionary["data"] = rawData
    rawDatasetDictionary["backgroundColor"] = backgroundColor

    rawDataset.push(rawDatasetDictionary)

    // Complete
    chartData["labels"] = rawLables
    chartData["datasets"] = rawDataset
    console.log(chartData);

    var ctx = document.getElementById("myChart");
    var myDoughnutChart = new Chart(ctx, {
        type: 'doughnut',
        data: chartData,
    });
}

function getRandomColor() {
    var letters = '0123456789ABCDEF'.split('');
    var color = '#';
    for (var i = 0; i < 6; i++ ) {
        color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
}
