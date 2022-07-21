$(function () {
  console.log("vm", vm)
  document.getElementById("dashboard").className += " active";
  if (vm.LoginType == "superadmin") {
    console.log("vm.AgentData", vm.AgentData)
    if (vm.AgentData != null) {
      document.getElementById("agentPassword").innerText = vm.AgentData.length;
    } else {
      document.getElementById("agentPassword").innerText = 0;
    }
    document.getElementById("agentPasswordDiv").style.display = "block";
  }
  if (vm.LoginType == "superadmin") {
    document.getElementById("notificationCount").innerText = vm.Count;
    document.getElementById("notificationDiv").style.display = "block";

  }
  if (vm.LoginType == "superadmin" || vm.LoginType == "staff") {

    var sel = document.getElementById('projectSelection');
    var opt = null;

    for (i = 0; i < vm.IDArray.length; i++) {
      opt = document.createElement('option');
      opt.value = vm.IDArray[i];
      opt.innerHTML = vm.Project[i];
      sel.appendChild(opt);
    }
    document.getElementById("projectData").style.display = "block";
    document.getElementById("projectSelection").value = vm.SessionProject
  }
  $("#agentButton").click(function () {
    $('#agentDropdown').empty();
    var strng = ""
    for (var i = 0; i < vm.AgentData.length; i++) {
      strng = strng + '<li>' + vm.AgentData[i][1] + " " + vm.AgentData[i][0] + " password changed" + '<li>' + '<br>'
    }
    $("#agentDropdown").append(strng)

  });
  var d = new Date();
  var currentYear = d.getFullYear();
  var currentMonth = d.getMonth() + 1;
  document.getElementById("yearpicker").value = currentYear;
  // document.getElementById("month").value = currentMonth;

  var arragedTotalBusCount = [];
  var month, year, weeks, monthNbr;
  $("#yearpicker").datepicker({
    onSelect: function () {
      $(this).change();
    },
    format: "yyyy",
    viewMode: "years",
    minViewMode: "years",
    autoclose: true,

  }).on("change", function () {
    $('#cover-spin').show(0);
    year = $(this).val();
    month = "";
    document.getElementById("month").value = "";
    weeks = "";
    document.getElementById("weeks").value = "";
    if (month == "") {
      month1 = currentMonth
    }
    var start = moment([parseInt(year), parseInt(month1) - 1]);
    var startDate = moment(start).startOf('month').toDate();
    var end = moment(start).endOf('month').toDate();
    $('#weeks').datepicker('setStartDate', startDate);
    $('#weeks').datepicker('setEndDate', end);
    $.ajax({
      type: 'POST',
      url: '/dashboardfilter',
      data: {
        'year': year,
        'month': month,
      },
      success: function (data) {
        var jsonData = JSON.parse(data);
        $('#cover-spin').hide(0);
        // totaTransChart.destroy(); 
        // livebusChart.destroy(); 
        // cardTrans.destroy(); 
        // chart.destroy(); 

        DrawAssociationGraph(jsonData[0], jsonData[10], arragedTotalBusCount)
        DrawMixedChartForTotalTransaction(jsonData[1], jsonData[2], jsonData[3])
        DrawMixedChartForCard(jsonData[4], jsonData[5], jsonData[6])
        DrawColumnBusLive(jsonData[7], jsonData[8])
        document.getElementById("tile1").innerText = jsonData[9][0]
        document.getElementById("tile2").innerText = "₹" + jsonData[9][1]
        document.getElementById("tile3").innerText = jsonData[9][2]
        document.getElementById("tile4").innerText = jsonData[9][3]
        document.getElementById("tile5").innerText = jsonData[9][4]
        document.getElementById("tile6").innerText = jsonData[9][5]
        document.getElementById("tile7").innerText = "₹" + jsonData[9][6]
        document.getElementById("tile8").innerText = "₹" + jsonData[9][8];
        document.getElementById("tile10").innerText = jsonData[9][7];
        document.getElementById("tile9").innerText = jsonData[9][9];

      }
    });
  });

  $('#month').datepicker({
    onSelect: function () {
      $(this).change();
    },
    format: 'MM',
    minViewMode: 'months',
    maxViewMode: 'months',
    startView: 'months',
    autoclose: true
  }).on("change", function () {
    $('#cover-spin').show(0);
    document.getElementById("weeks").value = ""
    month = $(this).val();
    if (year == null) {
      year = currentYear
    }
    monthNbr = monthNameToNum(month)
    var start = moment([parseInt(year), parseInt(monthNbr) - 1]);
    var startDate = moment(start).startOf('month').toDate();
    var end = moment(start).endOf('month').toDate();
    $('#weeks').datepicker('setStartDate', startDate);
    $('#weeks').datepicker('setEndDate', end);


    $.ajax({
      type: 'POST',
      url: '/dashboardfilter',
      data: {
        'year': year,
        'month': monthNbr,
      },
      success: function (data) {
        $('#cover-spin').hide(0);
        var jsonData = JSON.parse(data);
        // chart.render();
        DrawAssociationGraph(jsonData[0], jsonData[10], arragedTotalBusCount)
        DrawMixedChartForTotalTransaction(jsonData[1], jsonData[2], jsonData[3])
        DrawMixedChartForCard(jsonData[4], jsonData[5], jsonData[6])
        var liveStaus = [];
        if (jsonData[7] != null) {
          for (var i = 0; i < jsonData[7].length; i++) {
            var temp = { x: new Date(jsonData[5]), y: jsonData[7], indexLabel: jsonData[7][i].toString() }
            liveStaus.push(temp)
          }
        }
        DrawColumnBusLive(jsonData[7], jsonData[8])
        document.getElementById("tile1").innerText = jsonData[9][0]
        document.getElementById("tile2").innerText = "₹" + jsonData[9][1]
        document.getElementById("tile3").innerText = jsonData[9][2]
        document.getElementById("tile4").innerText = jsonData[9][3]
        document.getElementById("tile5").innerText = jsonData[9][4]
        document.getElementById("tile6").innerText = jsonData[9][5]
        document.getElementById("tile7").innerText = "₹" + jsonData[9][6];
        document.getElementById("tile8").innerText = "₹" + jsonData[9][8];
        document.getElementById("tile10").innerText = jsonData[9][7];
        document.getElementById("tile9").innerText = jsonData[9][9];

      }
    });
  });

  $("#weeks").datepicker({
    onSelect: function () {
      $(this).change();
    },
    changeMonth: false,
    changeYear: false,
    format: 'dd',
    multidate: true,
    startDate: moment().startOf('month').toDate(),
    endDate: moment().endOf('month').toDate(),

  }).on("change", function () {
    if (year == null) {
      year = currentYear
    }
    if (monthNbr == null) {
      monthNbr = currentMonth
    }
    $('#cover-spin').show(0);
    weeks = $(this).val();
    if (weeks != "") {
      $.ajax({
        type: 'POST',
        url: '/dashboardfilter',
        data: {
          'year': year,
          'month': monthNbr,
          'weeks': weeks,
        },
        success: function (data) {
          $('#cover-spin').hide(0);
          var jsonData = JSON.parse(data);
          DrawAssociationGraph(jsonData[0], jsonData[10], arragedTotalBusCount)
          DrawMixedChartForTotalTransaction(jsonData[1], jsonData[2], jsonData[3])
          DrawMixedChartForCard(jsonData[4], jsonData[5], jsonData[6])
          DrawColumnBusLive(jsonData[7], jsonData[8])

          document.getElementById("tile1").innerText = jsonData[9][0]
          document.getElementById("tile2").innerText = "₹" + jsonData[9][1]
          document.getElementById("tile3").innerText = jsonData[9][2]
          document.getElementById("tile4").innerText = jsonData[9][3]
          document.getElementById("tile5").innerText = jsonData[9][4]
          document.getElementById("tile6").innerText = jsonData[9][5]
          document.getElementById("tile7").innerText = "₹" + jsonData[9][6]
          document.getElementById("tile8").innerText = "₹" + jsonData[9][8];
          document.getElementById("tile10").innerText = jsonData[9][7]
          document.getElementById("tile9").innerText = jsonData[9][9]

        }
      });
    }
  });


  document.getElementById("dailyBus").innerHTML = vm.CountBus
  document.getElementById("tile1").innerText = vm.TileValues[0]
  document.getElementById("tile2").innerText = "₹" + vm.TileValues[1]
  document.getElementById("tile3").innerText = vm.TileValues[2]
  document.getElementById("tile4").innerText = vm.TileValues[3]
  document.getElementById("tile5").innerText = vm.TileValues[4]
  document.getElementById("tile6").innerText = vm.TileValues[5]
  document.getElementById("tile7").innerText = "₹" + vm.TileValues[6]
  document.getElementById("tile9").innerText = vm.TileValues[9];
  document.getElementById("tile8").innerText = "₹" + vm.TileValues[8];
  document.getElementById("tile10").innerText = vm.TileValues[7]
  // arrage totalbus count array like same as union chart
  if (vm.TotalBusName.indexOf("KOCHIWHEELZ") > -1) {
    arragedTotalBusCount[0] = vm.TotalBusArry[vm.TotalBusName.indexOf("KOCHIWHEELZ")];
  } else {
    arragedTotalBusCount[0] = 0;
  }
  if (vm.TotalBusName.indexOf("MY METRO") > -1) {
    arragedTotalBusCount[1] = vm.TotalBusArry[vm.TotalBusName.indexOf("MY METRO")];
  } else {
    arragedTotalBusCount[1] = 0;
  }
  if (vm.TotalBusName.indexOf("PRATEEKSHA") > -1) {
    arragedTotalBusCount[2] = vm.TotalBusArry[vm.TotalBusName.indexOf("PRATEEKSHA")];
  } else {
    arragedTotalBusCount[2] = 0;
  }

  if (vm.TotalBusName.indexOf("KMTC") > -1) {
    arragedTotalBusCount[3] = vm.TotalBusArry[vm.TotalBusName.indexOf("KMTC")];
  } else {
    arragedTotalBusCount[3] = 0;
  }
  if (vm.TotalBusName.indexOf("PBMS") > -1) {
    arragedTotalBusCount[4] = vm.TotalBusArry[vm.TotalBusName.indexOf("PBMS")];
  } else {
    arragedTotalBusCount[4] = 0;
  }
  if (vm.TotalBusName.indexOf("GCBT") > -1) {
    arragedTotalBusCount[5] = vm.TotalBusArry[vm.TotalBusName.indexOf("GCBT")];
  } else {
    arragedTotalBusCount[5] = 0;
  }
  if (vm.TotalBusName.indexOf("MUZIRIS") > -1) {
    arragedTotalBusCount[6] = vm.TotalBusArry[vm.TotalBusName.indexOf("MUZIRIS")];
  } else {
    arragedTotalBusCount[6] = 0;
  }
  console.log("arragedTotalBusCount", arragedTotalBusCount)
  //for barchart representation for onboreded bus with association
  DrawAssociationGraph(vm.BusCountArry, vm.AssociationArry, arragedTotalBusCount)

  //column chart and line chart combination for total transaction amount and ticket
  DrawMixedChartForTotalTransaction(vm.TotalTicketCount, vm.TransactionDate, vm.TotalTransactionAmount)

  //column chart and line chart combination for total card amount and ticket
  if (vm.TotalCardCount != null) {
    DrawMixedChartForCard(vm.TotalCardCount, vm.CardDate, vm.TotalCardAmount)

  }
  //column chart for live buses 

  DrawColumnBusLive(vm.LiveBusArray, vm.LiveBUsDate)
});
function DrawMixedChartForTotalTransaction(TotalTicketCount, TransactionDate, TotalTransactionAmount) {
  var totalTxnAmount = [];
  if (TotalTicketCount != null) {
    for (var i = 0; i < TotalTicketCount.length; i++) {
      var index1 = TotalTransactionAmount[i].toString()
      var temp = { y: TotalTransactionAmount[i], label: TransactionDate[i], indexLabel: "₹" + index1 }

      totalTxnAmount.push(temp)
    }
  }
  var totalTxnCount = [];
  if (TotalTicketCount != null) {
    for (var i = 0; i < TotalTicketCount.length; i++) {
      var temp = { y: TotalTicketCount[i], label: TransactionDate[i], indexLabel: TotalTicketCount[i].toString() }
      totalTxnCount.push(temp)
    }
  }

  var totaTransChart = new CanvasJS.Chart("totalTransactionChart", {
    animationEnabled: true,
    theme: "light2",
    title: {
      text: "Total Transaction Count and Amount",
      borderColor: "#b35900",
      borderThickness: 1,
      padding: 3,
      fontSize: 20
    },
    axisX: {
      title: "",
      //valueFormatString: "DD-MMM",
      labelAngle: -90,
      gridThickness: 0,
      lineThickness: 1,

    },
    axisY: [
      {
        title: "Millions",
        prefix: "₹",
        gridThickness: 0,
        lineThickness: 0,
        //labelFormatter: addSymbols
      },
      {
        title: "Transaction Amount"
      }
    ],
    axisY2: [
      {
        title: "Thousands",
        gridThickness: 0
      },
      {
        title: "Transaction Count",
      }
    ],


    toolTip: {
      shared: false
    },
    legend: {
      cursor: "pointer",
      //itemclick: toggleDataSeries
    },
    dataPointWidth: 10,
    data: [
      {
        type: "column",
        color: "#fee59b",
        axisYType: "secondary",
        name: "Sum of total txn count",
        showInLegend: true,
        // xValueFormatString: "DD-MMM",
        //yValueFormatString: "$#,##0",
        indexLabelPlacement: "inside",
        indexLabelFontColor: "black",
        dataPoints: totalTxnCount
      },
      {
        type: "line",
        lineColor: "#db8e4e",

        axisYIndex: 10,
        // interval: 10,
        name: "Sum of total txn amount",
        showInLegend: true,
        indexLabelFontColor: "black",
        dataPoints: totalTxnAmount
      },
    ]
  });
  //totaTransChart.destroy();
  totaTransChart.render();

  // $("#cardTransChart").CanvasJSChart(options);
}
function DrawMixedChartForCard(TotalCardCount, CardDate, TotalCardAmount) {

  var totalCardAmount = [];
  if (TotalCardCount != null) {
    for (var i = 0; i < TotalCardCount.length; i++) {
      var index1 = TotalCardAmount[i].toString()
      var temp = { y: TotalCardAmount[i], label: CardDate[i], indexLabel: "₹" + index1 }

      totalCardAmount.push(temp)
    }
  }
  var totalCardCount = [];
  if (TotalCardCount != null) {
    for (var i = 0; i < TotalCardCount.length; i++) {
      var temp = { y: TotalCardCount[i], label: CardDate[i], indexLabel: TotalCardCount[i].toString() }
      totalCardCount.push(temp)
    }
  }

  var cardTrans = new CanvasJS.Chart("totalCardChart", {

    //var options = {
    animationEnabled: true,
    theme: "light2",
    title: {
      borderColor: "#b35900",
      text: "Card Transaction Count and Amount",
      borderThickness: 1,
      padding: 3,
      fontSize: 20
    },
    axisX: {
      title: "",
      //valueFormatString: "DD-MMM",
      labelAngle: -90,
      gridThickness: 0,
      lineThickness: 1,

    },
    axisY: [
      {
        // title: "Millions",
        prefix: "₹",
        gridThickness: 0,
        lineThickness: 0,
        //labelFormatter: addSymbols
      },
      {
        title: "Card Transaction Amount"
      }
    ],
    axisY2: [
      {
        // title: "Thousands",
        gridThickness: 0
      },
      {
        title: "Card Transaction Count",
      }
    ],
    toolTip: {
      shared: false
    },
    legend: {
      cursor: "pointer",
      //itemclick: toggleDataSeries
    },
    dataPointWidth: 10,
    data: [
      {
        type: "column",
        color: "#fee59b",
        axisYType: "secondary",
        name: "Sum of total card txn count",
        showInLegend: true,
        // xValueFormatString: "DD-MMM",
        //yValueFormatString: "$#,##0",
        indexLabelPlacement: "inside",
        indexLabelFontColor: "black",
        dataPoints: totalCardCount
      },
      {
        type: "line",
        lineColor: "#db8e4e",
        // axisYIndex: 10,
        // interval: 10,
        name: "Sum of total card txn amount",
        indexLabelFontColor: "black",
        showInLegend: true,
        dataPoints: totalCardAmount
      },
    ]
  });
  cardTrans.render();

  // $("#totalCardChart").CanvasJSChart(options);
}

function DrawColumnBusLive(LiveBusArray, CardDate) {

  var liveStaus = [];
  if (LiveBusArray != null) {
    for (var i = 0; i < LiveBusArray.length; i++) {
      // var temp = { x: new Date(CardDate[i]), y: LiveBusArray[i],label: CardDate[i], indexLabel: LiveBusArray[i].toString() }
      var temp = { y: LiveBusArray[i], label: CardDate[i], indexLabel: LiveBusArray[i].toString() }

      liveStaus.push(temp)
    }
  }

  var livebusChart = new CanvasJS.Chart("liveBusChart", {

    animationEnabled: true,
    theme: "light2",
    title: {
      text: "Live Bus Status",
      borderColor: "#b35900",
      borderThickness: 1,
      padding: 3,
      fontSize: 20
    },
    axisX: {
      title: "",
      //valueFormatString: "MMM",
      labelAngle: -90,
      gridThickness: 0,
      lineThickness: 1,

    },
    axisY: [
      {
        gridThickness: 1,
        lineThickness: 1,
        //labelFormatter: addSymbols
      },
      {
        title: ""
      }
    ],

    toolTip: {
      shared: false
    },
    legend: {
      cursor: "pointer",
      //itemclick: toggleDataSeries
    },
    dataPointWidth: 20,
    data: [
      {
        type: "column",
        color: "#f3b183",
        indexLabelPlacement: "outside",
        indexLabelBackgroundColor: "#cae4b6",
        indexLabelFontColor: "black",
        indexLabelFontSize: 15,
        dataPoints: liveStaus
      },

    ]
  });

  livebusChart.render();
}
function addSymbols(e) {
  var suffixes = ["", "K", "M", "B"];
  var order = Math.max(Math.floor(Math.log(e.value) / Math.log(1000)), 0);

  if (order > suffixes.length - 1)
    order = suffixes.length - 1;

  var suffix = suffixes[order];
  return CanvasJS.formatNumber(e.value / Math.pow(1000, order)) + suffix;
}

function toggleDataSeries(e) {
  if (typeof (e.dataSeries.visible) === "undefined" || e.dataSeries.visible) {
    e.dataSeries.visible = false;
  } else {
    e.dataSeries.visible = true;
  }
  e.chart.render();
}
function DrawAssociationGraph(CountArr, associationName, totalBus) {
  console.log("CountArr", totalBus)
  if (associationName != null) {
    var arrayIndex = [];
    if (associationName.indexOf("KOCHIWHEELZ") > -1) {
      arrayIndex[0] = CountArr[associationName.indexOf("KOCHIWHEELZ")];
    } else {
      arrayIndex[0] = 0;
    }
    if (associationName.indexOf("MY METRO") > -1) {
      arrayIndex[1] = CountArr[associationName.indexOf("MY METRO")];
    } else {
      arrayIndex[1] = 0;
    }
    if (associationName.indexOf("PRATEEKSHA") > -1) {
      arrayIndex[2] = CountArr[associationName.indexOf("PRATEEKSHA")];
    } else {
      arrayIndex[2] = 0;
    }

    if (associationName.indexOf("KMTC") > -1) {
      arrayIndex[3] = CountArr[associationName.indexOf("KMTC")];
    } else {
      arrayIndex[3] = 0;
    }
    if (associationName.indexOf("PBMS") > -1) {
      arrayIndex[4] = CountArr[associationName.indexOf("PBMS")];
    } else {
      arrayIndex[4] = 0;
    }
    if (associationName.indexOf("GCBT") > -1) {
      arrayIndex[5] = CountArr[associationName.indexOf("GCBT")];
    } else {
      arrayIndex[5] = 0;
    }
    if (associationName.indexOf("MUZIRIS") > -1) {
      arrayIndex[6] = CountArr[associationName.indexOf("MUZIRIS")];
    } else {
      arrayIndex[6] = 0;
    }
    console.log("arrayIndex", arrayIndex)
  }
  var onboardedBus = [];
  //var color = ["#000099", "#e65c00", "#808080", "#00e673", "#ffb31a", "#9B32DB", "#DB326F"]
  if (arrayIndex == null) {
    var arrayIndex = [];
    arrayIndex[0] = 0;
    arrayIndex[1] = 0
    arrayIndex[2] = 0
    arrayIndex[3] = 0
    arrayIndex[4] = 0
    arrayIndex[5] = 0
    arrayIndex[6] = 0
  }
  for (var i = 0; i < arrayIndex.length; i++) {
    var temp = { y: arrayIndex[i], label: i + 1, indexLabel: arrayIndex[i].toString(), color: "#808080" }
    onboardedBus.push(temp)
  }


  var onboardedTotalBus = [];
  if (totalBus != null) {
    for (var i = 0; i < totalBus.length; i++) {
      var temp = { y: totalBus[i], label: i + 1, indexLabel: totalBus[i].toString(), color: "#e65c00" }
      onboardedTotalBus.push(temp)
    }
  }
  var chart = new CanvasJS.Chart("barchart", {
    //var options = {
    // axisX:{
    //   gridThickness: 0,
    //   tickLength: 0,
    //   lineThickness: 0,
    //   labelFontColor: "transparent"
    // },
    axisY: {
      minimum: -1
    },
    animationEnabled: true,
    theme: "light2",
    title: {
      display: true,
      text: 'Union wise Onboarded Buses',
      borderColor: "#b35900",
      borderThickness: 1,
      padding: 3,
      fontSize: 20
    },

    //dataPointWidth: 25,
    //color: ["#0052cc", "#e65c00", "#808080", "#ffb31a", "#3385ff", "#00e673", "#000099"],
    data: [
      {
        type: "column",
        indexLabelPlacement: "outside",
        indexLabelFontColor: "black",
        dataPoints: onboardedTotalBus
      },
      {
        type: "column",
        indexLabelPlacement: "outside",
        indexLabelFontColor: "black",
        dataPoints: onboardedBus
      }

    ]

  });
  chart.render();
}
var months = [
  'January', 'February', 'March', 'April', 'May',
  'June', 'July', 'August', 'September',
  'October', 'November', 'December'
];
function monthNameToNum(monthname) {
  var month = months.indexOf(monthname);
  return month ? month + 1 : 0;
}

$("#projectSelection").change(function () {
  $('#cover-spin').show(0);
  $.ajax({
    type: 'POST',
    url: "/sessionProject",
    data: {
      'projectSelection': $("#projectSelection").val(),
      'projectSelectionName': $("#projectSelection option:selected").text()
    },
    success: function (data) {
      var jsonData = JSON.parse(data);
      console.log(jsonData[0]);
      if (jsonData[0] == "true") {
        window.location = "/dashboard";
      } else if (jsonData[0] == "service" || jsonData[0] == "paycraft") {
        window.location = "/detailed_dashboard";
      } else if (jsonData[0] == "association") {
        window.location = "/associationDashboard";
      } else {
        $("#login_err").css({ "color": "red", "font-size": "15px", "margin-left": "33px" });
        $("#login_err").html("Invalid Username or Password!");
      }
    }
  });
});