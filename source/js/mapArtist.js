//token
mapboxgl.accessToken = 'pk.eyJ1IjoibGVvYnJ2IiwiYSI6ImNsZmJmcms3MjBjdXkzcG12cmc5Nno1bmwifQ.Mbvg3DAfN_1-gscArP8hFg';

//create map
const map = new mapboxgl.Map({
    container: 'map',
    style: 'mapbox://styles/mapbox/streets-v11',
    center: [-96, 37.8],
    zoom: 3
});

function addMarker(cityName, popupText) {
    //send a request to find the longitude and latitude of the city
    const url = 'https://api.mapbox.com/geocoding/v5/mapbox.places/' + cityName + '.json?access_token=' + mapboxgl.accessToken;
    fetch(url)
        .then(response => response.json())
        .then(data => {
            //we retrieve the geographical coordinates of the first answer
            const longitude = data.features[0].center[0];
            const latitude = data.features[0].center[1];

            //add a point on the map to the coordinates of the city
            const marker = new mapboxgl.Marker()
                .setLngLat([longitude, latitude])
                .addTo(map);

            //add a popup dates to point
            const popup = new mapboxgl.Popup({ offset: 25 })
                .setText(popupText)
                .addTo(map);

            marker.setPopup(popup);
        });
}