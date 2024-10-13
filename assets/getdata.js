const cactusSocket = new WebSocket(
    `ws://${location.host}/cactus`,
);

cactusSocket.onerror = (event) =>{
    console.log("WebSocket error: ", event);
}

cactusSocket.onmessage = (event) => {
    /**
     * @typedef Cactus
     * @type {object}
     * @property {string} img
     * @property {number} x
     * @property {number} y
     */

    /**
     * @type Cactus[]
     */
    const cactus = JSON.parse(event.data);
    cactus.forEach((item) => {
        const myImage = new Image(16, 16);
        myImage.src = item.img;
        myImage.style.position = "absolute"
        myImage.style.top = item.y + "px"
        myImage.style.left = item.x + "px"
        document.body.appendChild(myImage);
    })
}

const mapSocket = new WebSocket(
    `ws://${location.host}/map`,
);

/**
 * @typedef Anomaly
 * @type {object}
 * @property {string} id
 * @property {number} r1
 * @property {number} r2
 * @property {HTMLImageElement} image
 * @property {HTMLImageElement} image2
 * @property {number} x
 * @property {number} y
 */

/**
 * @type Anomaly[]
 */
let anomalies = [];


let bounties = [];

mapSocket.onmessage = (event) => {

    const map = JSON.parse(event.data);
    anomalies = anomalies.filter(
        (item) => {
            let notFound = true
            map.anomalies.forEach(
                (anomaly) => {
                    if (item.id === anomaly.id) {
                        notFound = false
                    }
                }
            )

            if (notFound) {
                item.image.parentElement.removeChild(item.image)
                item.image2.parentElement.removeChild(item.image2)
            }

            return !notFound;
        }
    )

    map.anomalies.forEach(
        (item) => {
            let index = -1
            anomalies.forEach(
                (anomaly, i) => {
                    if (item.id === anomaly.id){
                        index = i
                    }
                }
            )

            if (index < 0) {
                const imgOuter = new Image(item.effectiveRadius * 2, item.effectiveRadius * 2);
                imgOuter.src = "anomalyborder.png"
                imgOuter.style.position = "absolute"
                imgOuter.style.top = item.y - item.effectiveRadius + "px"
                imgOuter.style.left = item.x - item.effectiveRadius + "px"
                document.body.appendChild(imgOuter);

                const imgInner = new Image(item.radius * 2, item.radius * 2);
                if (item.strength > 0) {
                    imgInner.src = "dangerin.png"
                } else {
                    imgInner.src = "dangerout.png"
                }
                imgInner.style.position = "absolute"
                imgInner.style.top = item.y-item.radius + "px"
                imgInner.style.left = item.x-item.radius + "px"
                document.body.appendChild(imgInner);

                let anomaly = {
                    id: item.id,
                    r1: item.radius,
                    r2: item.effectiveRadius,
                    image: imgInner,
                    image2: imgOuter,
                    x: item.x,
                    y: item.y
                }

                anomalies.push(anomaly)
            } else {
                anomalies[index].image.style.top = item.y - item.radius + "px"
                anomalies[index].image.style.left = item.x - item.radius + "px"
                anomalies[index].image2.style.top = item.y - item.effectiveRadius + "px"
                anomalies[index].image2.style.left = item.x - item.effectiveRadius + "px"
            }
        }
    )

    bounties = bounties.filter(
        (item) => {
            let notFound = true
            map.bounties.forEach(
                (bounty) => {
                    if (item.x === bounty.x && item.y === bounty.y) {
                        notFound = false
                    }
                }
            )

            if (notFound) {
                item.image.parentElement.removeChild(item.image)
            }

            return !notFound;
        }
    )

    map.bounties.forEach(
        (item) => {
            let index = -1
            bounties.forEach(
                (bounty, i) => {
                    if (item.x === bounty.x && item.y === bounty.y) {
                        index = i
                    }
                }
            )

            if (index < 0) {
                const img = new Image(item.radius * 2, item.radius * 2);
                img.src = "bounty.png"
                img.style.position = "absolute"
                img.style.top = item.y - item.radius + "px"
                img.style.left = item.x - item.radius + "px"
                document.body.appendChild(img);

                let bounty = {
                    x: item.x,
                    y: item.y,
                    image: img,
                }

                bounties.push(bounty)
            }
        }
    )
}