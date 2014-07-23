$( document ).ready(function() {
  $('.title').click(function (e) {
    e.preventDefault();
    $(this).next().slideToggle();
  });

});