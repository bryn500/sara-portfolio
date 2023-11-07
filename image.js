// using this until cross platform go implementation of libvips
const sharp = require("sharp");

// Input JPG file path
const imageTransformMap = [
  { input: "input.jpg", output: "output.webp" },
  { input: "input.jpg", output: "output2.webp" },
];

transformImage = (item) => {
  console.log(`converting: ${item.output}`);
  return sharp(item.input).webp({ quality: 85 }).toFile(item.output);
};

Promise.all(imageTransformMap.map(transformImage))
  .then(() => {
    console.log("All transformations are done");
  })
  .catch((err) => {
    console.error(err);
  });
