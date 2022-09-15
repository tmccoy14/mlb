// Project Name Validation Function
$(document).ready(function () {
  $("#myform").on("submit", function () {
    var listValue = document.getElementById("listValue").value;

    if (listValue === "Players") {
      var playerValue = document.getElementById("player_name").value;

      let file = "http://localhost:1323/lookup/players?name=" + playerValue
      fetch (file)
      .then(x => x.text())
      .then(y => {
        document.getElementById("allow").innerHTML = y;
      });

      $("#pageloader").fadeIn();

      return false;
    } else if (listValue === "Stats") {
      var playerValue = document.getElementById("player_name").value;
      var seasonValue = document.getElementById("season").value;

      let file = "http://localhost:1323/lookup/stats?name=" + playerValue + "&season=" + seasonValue
      fetch (file)
      .then(x => x.text())
      .then(y => {
        document.getElementById("allow").innerHTML = y;
      });

      $("#pageloader").fadeIn();

      return false;
    } else if (listValue === "Teams") {
      var playerValue = document.getElementById("player_name").value;
      var seasonValue = document.getElementById("season").value;

      let file = "http://localhost:1323/lookup/teams?name=" + playerValue + "&season=" + seasonValue
      fetch (file)
      .then(x => x.text())
      .then(y => {
        document.getElementById("allow").innerHTML = y;
      });

      $("#pageloader").fadeIn();

      return false;
    }
  });
});

// Reset All On Click New Search
$(document).ready(function () {
  $("#clear").on("click", function () {
    document.getElementById("myform").reset();
    document.getElementById("allow").innerHTML = "";
    $("input[type=submit]").hide();
    $("input[id=player_name]").hide();
    $("input[id=season]").hide();
    $("label").hide();
    $("#pageloader").hide();

  });
});

// Show Options On Create New Project Otherwise Hide All
$(document).ready(function() {
  $("select")
    .change(function() {
      $(this)
        .find("option:selected")
        .each(function() {
          var optionValue = $(this).attr("value");
          if (optionValue === "Select Lookup Type") {
            $("input[type=submit]").hide();
            $("input[id=player_name]").hide();
            $("input[id=season]").hide();
            $("label[id=player_label]").hide();
            $("label[id=season_label]").hide();
          } else if (optionValue === "Players") {
            document.getElementById("player_name").required = true;
            document.getElementById("season").required = false;
            $(".box")
              .not("." + optionValue)
              .hide();
            $("." + optionValue).show();
            $("input[id=player_name]").show();
            $("label[id=player_label]").show();
            $("input[id=season]").hide();
            $("label[id=season_label]").hide();
            $("input[type=submit]").show();
          } else {
            document.getElementById("player_name").required = true;
            document.getElementById("season").required = true;
            $(".box").hide();
            $("input[id=player_name]").show();
            $("input[id=season]").show();
            $("label").show();
            $("input[type=submit]").show();
          }
        });
    })
    .change();
});

// On Page Load Get List Of Existing Projects
$(document).ready(function() {
  $("input[type=submit]").hide();
  $("input[id=player_name]").hide();
  $("input[id=season]").hide();
  $("label").hide();

  onload = () => {
    var projects = '["Select Lookup Type", "Players", "Teams", "Stats"]';

    var values = JSON.parse(projects);

    var list1 = document.getElementById("listValue");

    var i;
    for (i = 0; i < values.length; i++) {
      list1.options[i] = new Option(values[i], values[i]);
      }
  };
});
