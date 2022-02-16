<template>
  <div style="min-width: 540px;width:600px;">
    <div class="eleme">
      <el-upload
        class="upload-demo"
        ref="upload"
        action="https://jsonplaceholder.typicode.com/posts/"
        :before-upload="beforeUpload"
        :on-preview="handlePreview"
        :on-remove="handleRemove"
        :auto-upload="true"
        :show-file-list="false"
      >
        <el-button slot="trigger" size="small" type="primary">选择图片</el-button>
        <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload">上传头像</el-button>
      </el-upload>
    </div>
    <div>
      <br />
      <el-button type="primary" icon="el-icon-refresh-right" circle @click="rotateRight"></el-button>
      <el-button type="success" icon="el-icon-refresh-left" circle @click="rotateLeft"></el-button>
      <el-button type="danger" icon="el-icon-plus" circle @click="changeScale(1)"></el-button>
      <el-button type="warning" icon="el-icon-minus" circle @click="changeScale(-1)"></el-button>
    </div>
    <div class="cropper">
      <div class="cropper-content" style="margin-top:60px;margin-left:60px;">
        <div class="cropper">
          <vueCropper
            ref="cropper"
            :img="option.img"
            :outputSize="option.size"
            :outputType="option.outputType"
            :info="true"
            :full="option.full"
            :canMove="option.canMove"
            :canMoveBox="option.canMoveBox"
            :original="option.original"
            :autoCrop="option.autoCrop"
            :autoCropWidth="option.autoCropWidth"
            :autoCropHeight="option.autoCropHeight"
            :fixedBox="option.fixedBox"
            @realTime="realTime"
            @imgLoad="imgLoad"
          ></vueCropper>
        </div>
        <div style="margin-left:20px;">
          <div
            class="show-preview"
            :style="{'width': '150px', 'height':'155px',  'overflow': 'hidden', 'margin': '5px'}"
          ></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      headImg: '',
      //剪切图片上传
      crap: false,
      previews: {},
      option: {
        img: '',
        outputSize: 1, //剪切后的图片质量（0.1-1）
        full: false, //输出原图比例截图 props名full
        outputType: 'png',
        canMove: true,
        original: false,
        canMoveBox: true,
        autoCrop: true,
        autoCropWidth: 150,
        autoCropHeight: 150,
        fixedBox: false
      },
      fileName: '', //本机文件地址
      downImg: '#',
      imgFile: '',
      uploadImgRelaPath: '' //上传后的图片的地址（不带服务器域名）
    }
  },

  methods: {
    submitUpload(file) {
      // this.$refs.upload.submit();
      this.finish('blob')
    },
    handleRemove(file, fileList) {
      console.log(file, fileList)
    },
    handlePreview(file) {
      console.log(file)
      //   let data = window.URL.createObjectURL(new Blob([file]));
      //   this.option.img = data;
    },
    beforeUpload(file) {
      console.log('上传文件')
      console.log(file)
      let data = window.URL.createObjectURL(new Blob([file]))
      this.fileName = file.name
      this.option.img = data
    },
    //放大/缩小
    changeScale(num) {
      console.log('changeScale')
      num = num || 1
      this.$refs.cropper.changeScale(num)
    },
    //坐旋转
    rotateLeft() {
      console.log('rotateLeft')
      this.$refs.cropper.rotateLeft()
    },
    //右旋转
    rotateRight() {
      console.log('rotateRight')
      this.$refs.cropper.rotateRight()
    },
    //上传图片（点击上传按钮）
    finish(type) {
      console.log('finish')
      let _this = this
      let formData = new FormData()
      // 输出
      if (type === 'blob') {
        this.$refs.cropper.getCropBlob(data => {
          let img = window.URL.createObjectURL(data)
          this.model = true
          this.modelSrc = img
          formData.append('file', data, this.fileName)
          this.$axios
            .post(config.upLoadFileURL, formData, {
              contentType: false,
              processData: false,
              headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
            })
            .then(response => {
              var res = response.data
              if (res == 'success') {
                console.log('上传成功！')
              }
            })
        })
      } else {
        this.$refs.cropper.getCropData(data => {
          this.model = true
          this.modelSrc = data
        })
      }
    },
    // 实时预览函数
    realTime(data) {
      console.log('realTime')
      this.previews = data
    },
    imgLoad(msg) {
      console.log('imgLoad')
      console.log(msg)
    }
  }
}
</script>
<style lang="scss">
.cropper-content {
  min-width: 540px;
  display: flex;
  .cropper {
    width: 260px;
    height: 260px;
  }
  .show-preview {
    flex: 1;
    -webkit-flex: 1;
    display: flex;
    display: -webkit-flex;
    justify-content: center;
    -webkit-justify-content: center;
    .preview {
      overflow: hidden;
      border-radius: 50%;
      border: 1px solid #cccccc;
      background: #cccccc;
      margin-left: 40px;
    }
  }
}
.cropper-content .show-preview .preview {
  margin-left: 0;
  -moz-user-select: none;
  -webkit-user-select: none;
  -ms-user-select: none;
  -khtml-user-select: none;
  user-select: none;
}
</style>