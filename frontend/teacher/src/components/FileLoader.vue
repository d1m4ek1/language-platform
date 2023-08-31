<script>
export default {
  name: "FileLoader",
  props: {
    idFor: String,
    typeFile: String,
    nameFormData: String,
    editPreview: ""
  },
  data() {
    return {
      preview: '',
    };
  },
  methods: {
    listenUploadFile(e) {
      const file = e.target.files[0];
      // const validFileTypeSize = this.validFile(file);

      // if (!validFileTypeSize.success) return this.$emit('onUploadFile', validFileTypeSize);

      this.$emit('onUploadFile', {
        file,
        nameFormData: this.nameFormData,
        success: true
      });

      this.preview = window.URL.createObjectURL(file);
    },
    validFile(file) {
      const parsedSize = this.formatBytes(file.size);

      if (file.type.startsWith('image/')) {
        if ((parsedSize.typeSize === 'Bytes' || parsedSize.typeSize === 'KB' || parsedSize.typeSize === 'MB') && file.size <= 5242880) {
          return { success: true };
        }

        return {
          message: 'Размер файла превышает 5 МБ.',
          success: false
        };
      }
      return {
        message: 'Не верный тип файла',
        success: false
      };
    },
    formatBytes(bytes, decimals = 2) {
      if (!+bytes) return '0 Bytes';

      const k = 1024;
      const dm = decimals < 0 ? 0 : decimals;
      const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

      const i = Math.floor(Math.log(bytes) / Math.log(k));

      return {
        parsedValue: parseFloat((bytes / Math.pow(k, i)).toFixed(dm)),
        typeSize: sizes[i]
      };
    },
    removeFile() {
      this.preview = '';
      this.$emit('onUploadFile', {
        remove: true,
        nameFormData: this.nameFormData
      });
    }
  },
  watch: {
    editPreview: {
      immediate: true,
      handler(n) {
        if (n !== undefined && n !== "") {
          this.preview = n
        }
      }
    }
  },
};
</script>

<template>
  <div class='file-loader-block'>
    <template v-if='typeFile === "image"'>
      <template v-if='preview !== ""'>
        <img :src='preview' alt='Предпросмотр изображения'>
        <button class='btn-red' @click='removeFile'>Удалить</button>
      </template>
      <template v-else>
        <input :id='idFor' accept='image/png, image/jpg, image/jpeg, image/gif' type='file' @change='listenUploadFile'>
        <label :for='idFor'>Загрузить изображение</label>
      </template>
    </template>
    <template v-if='typeFile === "video"'>
      <template v-if='preview !== ""'>
        <video controls>
          <source :src='preview'>
        </video>
        <button class='btn-red' @click='removeFile'>Удалить</button>
      </template>
      <template v-else>
        <input :id='idFor' accept='video/mp4' type='file' @change='listenUploadFile'>
        <label :for='idFor'>Загрузить видео</label>
      </template>
    </template>
    <template v-if='typeFile === "audio"'>
      <template v-if='preview !== ""'>
        <audio controls>
          <source :src='preview'>
        </audio>
        <button class='btn-red' @click='removeFile'>Удалить</button>
      </template>
      <template v-else>
        <input :id='idFor' accept='audio/mpeg' type='file' @change='listenUploadFile'>
        <label :for='idFor'>Загрузить аудио файл</label>
      </template>
    </template>
  </div>
</template>

<style scoped>
.file-loader-block {
  width: 100%;
  min-height: 200px;
  background-color: #a9d5ff;
  border-radius: 10px;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
}

.file-loader-block > button {
  position: absolute;
  right: 10px;
  top: 10px;
}

.file-loader-block > img, .file-loader-block > video {
  width: 100%;
}


.file-loader-block > input {
  opacity: 0;
  position: absolute;
  width: 100%;
  height: 100%;
}

.file-loader-block > label {
  display: block;
}

audio {
  width: 100%;
}
</style>