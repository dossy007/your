$(function () {
  $(".video").click(function () {
    var theModal = $(this).data("target"),
    videoSRC = $(this).attr("data-video"),
    videoSRCauto = videoSRC + "?modestbranding=1&rel=0&controls=0&showinfo=0&html5=1&autoplay=1";
    $(theModal + ' iframe').attr('src', videoSRCauto);
    $(theModal + ' button.close').click(function () {
      $(theModal + ' iframe').attr('src', '');
    });
  });
  $('#exampleModal').on('hidden.bs.modal', function () {
    var $this = $(this).find('iframe')
    $this.attr('src', "");
  });

  $('.head-image').slick({
    dots:true,
    infinite: true,
    speed: 1000,
    slidesToShow: 1,
    adaptiveHeight: true,
    autoplay: true,
    fade: true,
    pauseOnHover: false
  });
});