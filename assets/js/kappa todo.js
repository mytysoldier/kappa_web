$(function () {
  // TODO入力エリア入力時の動作
  $("#todo_text").on("input", function () {
    if ($("#todo_text").val()) {
      // テキストが入力されていればTODO追加ボタンを活性にする
      $("#add_todo").prop("disabled", false);
    } else {
      // テキスト未入力であればTODO追加ボタンを非活性に戻す
      $("#add_todo").prop("disabled", true);
    }
  });

  // TODO追加ボタン押下時の動作
  $("#add_todo").on("click", function () {
    // TODO追加リクエスト送信
    $.ajax({
      url: "/add_todo",
      type: "POST",
      data: {
        todo: $("#todo_text").val(),
      },
    })
      .done(function (data) {
        window.location.href = "/kappa_todo";
      })
      .fail(function (data) {
        window.alert("TODO追加失敗");
      });
  });

  // 未完了TODO左のチェックボックス押下時の動作
  $("#uncomp-todos input:checkbox").on("change", function (e) {
    var checked = false;
    $.each($("#uncomp-todos li div :checkbox"), function (index, element) {
      if ($(element).is(":checked")) {
        checked = true;
      }
    });
    if (checked) {
      $("#complete_todo").prop("disabled", false);
    } else {
      $("#complete_todo").prop("disabled", true);
    }
  });

  // 完了済みTODO左のチェックボックス押下時の動作
  $("#comp-todos input:checkbox").on("change", function (e) {
    var checked = false;
    $.each($("#comp-todos li div :checkbox"), function (index, element) {
      if ($(element).is(":checked")) {
        checked = true;
      }
    });
    if (checked) {
      $("#uncomplete_todo").prop("disabled", false);
    } else {
      $("#uncomplete_todo").prop("disabled", true);
    }
  });

  // 完了にするボタン押下時の動作
  $("#complete_todo").on("click", function () {
    var requests = [];
    $.each($("#uncomp-todos li div :checkbox"), function (index, element) {
      if ($(element).is(":checked")) {
        requests.push($(element).attr("data-todo"));
      }
    });
    // TODO完了更新リクエスト送信
    $.ajax({
      url: "/update_todo",
      type: "POST",
      data: {
        todos: requests.join(","),
        mode: "comp",
      },
    })
      .done(function (data) {
        window.location.href = "/kappa_todo";
      })
      .fail(function (data) {
        window.alert("TODO更新失敗");
      });
  });

  // 未完にするボタン押下時の動作
  $("#uncomplete_todo").on("click", function () {
    var requests = [];
    $.each($("#comp-todos li div :checkbox"), function (index, element) {
      if ($(element).is(":checked")) {
        requests.push($(element).attr("data-todo"));
      }
    });
    // TODO未完更新リクエスト送信
    $.ajax({
      url: "/update_todo",
      type: "POST",
      data: {
        todos: requests.join(","),
        mode: "uncomp",
      },
    })
      .done(function (data) {
        window.location.href = "/kappa_todo";
      })
      .fail(function (data) {
        window.alert("TODO更新失敗");
      });
  });
});
