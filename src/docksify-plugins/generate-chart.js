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


// Register chart generate on root document only after HTML has been appended to the DOM
// Docsify custom plugin boilerplate found https://docsify.js.org/#/write-a-plugin?id=template
(function () {
    var generateChartOnRoot = function (hook, vm) {
        // Invoked one time after rendering the initial page
        hook.ready(function () {
            if (window.location.pathname === '/' || window.location.pathname === '/index.html') {
                generateChart('chart-canvas');
            }
        });
    };

    // Add plugin to docsify's plugin array
    $docsify = $docsify || {};
    $docsify.plugins = [].concat(generateChartOnRoot, $docsify.plugins || []);
})();