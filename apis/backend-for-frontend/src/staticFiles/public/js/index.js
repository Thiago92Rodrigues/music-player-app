document.addEventListener('DOMContentLoaded', function () {
  initCarousel();
});

function initCarousel() {
  const carouselViewport = document.querySelector('.carousel .carousel__viewport');

  const slidePreviousButtons = document.querySelectorAll('.carousel .carousel__viewport .carousel__prev');
  const slideNextButtons = document.querySelectorAll('.carousel .carousel__viewport .carousel__next');
  const navigationButtons = document.querySelectorAll('.carousel .carousel__navigation .carousel__navigation__button');

  let currentSlide = 1;
  let okToMove = true;

  slidePreviousButtons.forEach(element => element.addEventListener('click', scrollHorizontally));
  slideNextButtons.forEach(element => element.addEventListener('click', scrollHorizontally));
  navigationButtons.forEach(element => element.addEventListener('click', scrollHorizontally));

  function scrollHorizontally(event) {
    event.preventDefault();

    let targetElement = event.target || event.srcElement;

    if (targetElement.tagName == 'svg') {
      targetElement = targetElement.parentNode;
    } else if (targetElement.tagName == 'path') {
      targetElement = targetElement.parentNode.parentNode;
    }

    const link = targetElement.href;
    const internalLink = link.split('#')[1];
    const slideNumber = internalLink[internalLink.length - 1];

    const destinationElement = document.querySelector(`#${internalLink}`);

    if (destinationElement) {
      const leftPosition = destinationElement.offsetLeft;

      carouselViewport.scrollLeft = leftPosition;

      currentSlide = slideNumber;
    }
  }

  carouselViewport.addEventListener('mouseenter', event => {
    okToMove = false;
  });

  carouselViewport.addEventListener('mouseleave', event => {
    okToMove = true;
  });

  window.setInterval(() => {
    if (okToMove) {
      slideNextButtons[currentSlide - 1].click();
    }
  }, 5000);
}
