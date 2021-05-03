module.exports = {
  transpileDependencies: ["vuetify"],
  devServer: {
    proxy: {
      "/memo/": {
        target: "http://localhost:8080",
      },
    },
  },
};
