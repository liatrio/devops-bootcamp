var bootcampMetadata = {}

function generateChart(canvasId) {
    var canvas = document.getElementById(canvasId);
    // var canvas = document.createElement('canvas');
    // container.appendChild(canvas);

    canvas.width = 400;
    canvas.height = 400;

    // Set the data for the chart
    var data = {
        labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
        datasets: [
            {
                label: 'My Dataset',
                data: [10, 20, 30, 40, 50, 60, 70],
                backgroundColor: 'rgba(255, 99, 132, 0.2)',
                borderColor: 'rgba(255, 99, 132, 1)',
                borderWidth: 1
            }
        ]
    };

    // Set the options for the chart
    var options = {
        scales: {
            y: {
                ticks: {
                    beginAtZero: true
                }
            }
        }
    };

    // Create a new bar chart
    var myChart = new Chart(canvas, {
        type: 'bar',
        data: data,
        options: options
    });
}

function generateWordCloud(canvasId) {
    var canvas = document.getElementById(canvasId);
    const ctx = canvas.getContext('2d');
    // This object will hold the number of occurances of a technology in the bootcamp
    let techCount = {};

    // Objects are not inherintly itterable in javascript so we must converte the values to an array.
    for (const doc of Object.values(bootcampMetadata)) {

        if ('technolgies' in doc) {
            // we have some top level technologies not associated to an exercise

        }
        // loop over exercises
        if ('exercises' in doc) {
            for (const exercise of Object.values(doc.exercises)) {
                if ('technologies' in exercise) {
                    for (const tech of exercise.technologies) {
                        if (tech in techCount) {
                            techCount[tech] += 1;
                        } else {
                            techCount[tech] = 1;
                        }
                    }
                }
            }
        }
    }

    const data = {
        labels: Object.keys(techCount),
        datasets: [
            {
                data: Object.values(techCount)
            }
        ]
    }

    const options = {
        plugins: {
            wordcloud: {
                fontStyle: 'normal',
                fontColor: '#000',
                fontFamily: 'Helvetica, Arial, sans-serif',
                minFontSize: 12,
                maxFontSize: 50
            }
        }
    };

    const myChart = new Chart(ctx, {
        type: 'wordCloud',
        data: data,
        options: options
    });

    console.log(techCount)
}


// Register chart generate on root document only after HTML has been appended to the DOM
// Docsify custom plugin boilerplate found https://docsify.js.org/#/write-a-plugin?id=template
(function () {
    var generateChartOnRoot = function (hook, vm) {
        // Invoked one time after rendering the initial page
        hook.ready(function () {
            if (window.location.pathname === '/' || window.location.pathname === '/index.html') {
                generateChart('chart-canvas');
                generateWordCloud('chart-canvas');
            }
        });
    };

    // Add plugin to docsify's plugin array
    $docsify = $docsify || {};
    $docsify.plugins = [].concat(generateChartOnRoot, $docsify.plugins || []);
})();