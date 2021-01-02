$(function () {
  // かっぱを召喚ボタン押下時の動作
  $("#add_kappa").on("click", function () {
    var count = $("#add_kappa_count").val();
    // 選択件数分追加
    for (var i = 0; i < count; i++) {
      $("#contents").append(
        '<div class="kappa_create col-2"><img src="/assets/image/youkai_kappa.png" class="img-thumbnail"></div>'
      );
    }
    // かっぱを召喚したため帰還ボタンを活性
    $("#return_kappa").prop("disabled", false);
    // かっぱの数を計算
    var all_kappa = Number($(".kappa_count").text()) + Number(count);
    $(".kappa_count").text(all_kappa);
  });
  // かっぱを帰還ボタン押下時の動作
  $("#return_kappa").on("click", function () {
    var count = $("#return_kappa_count").val();
    // 選択件数分最後の要素から順に削除
    for (var i = 0; i < count; i++) {
      $("#contents").children().last().remove();
    }
    if ($("#contents").children().length == 0) {
      // 帰還させるかっぱがいなくなったため帰還ボタンを非活性
      $("#return_kappa").prop("disabled", true);
    }
    // かっぱの数を計算
    var all_kappa = Number($(".kappa_count").text()) - Number(count);
    // マイナス値になった場合は、0にする
    if (all_kappa < 0) {
      all_kappa = 0;
    }
    $(".kappa_count").text(all_kappa);
  });
});
