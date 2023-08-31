import { defineStore } from 'pinia';

export const useClient = defineStore('client', {
  state: () => ({
    dataClient: {
      firstName: '',
      lastName: '',
      aboutMe: '',
      regDate: '',
      avatar: '',
      isTeacher: null,
      interfaceLang: 'ru',
      languageItems: [
        {
          'lang': 'en',
          'lvl': 'B1'
        }
      ]
    },
    isGetDataClient: false
  }),
  getters: {
    getFullName: (state) => `${state.dataClient.firstName} ${state.dataClient.lastName}`,
    getAvatar: (state) => state.dataClient.avatar || '/default/avatar/1.png',
    getTextIsTeacher: (state) => state.dataClient.isTeacher ? 'Преподаватель' : 'Ученик',
    getAboutMe: (state) => state.dataClient.aboutMe || 'Вы о себе ни чего не написали',
    getLanguageItems: (state) => {
      const langItems = state.dataClient.languageItems;
      if (langItems === null) return 'Вы не указали преподаваемые языки';

      return langItems.map(item => {
        let langWord = 'Английский';

        switch (item.lang) {
          case 'en':
            langWord = 'Английский';
            break;
        }

        return `${langWord} - ${item.lvl}`;
      });
    },
    getDate: (state) => {
      let date = new Date(state.dataClient.regDate);
      return date.toLocaleDateString('ru-RU', {
        weekday: 'long',
        year: 'numeric',
        month: 'long',
        day: 'numeric'
      });
    }
  },
  actions: {
    async sendGetDataClient() {
      try {
        await fetch('/teacher/profile/data-client')
          .then(response => response.json())
          .then(response => {
            const resp = response.dataClient;
            Object.keys(resp).forEach(key => {
              if (this.dataClient[key] !== undefined) {
                this.dataClient[key] = resp[key];
              }
            });

            this.isGetDataClient = true;
          });
      } catch (err) {
        console.error(err);
      }
    },
    async sandSaveDataClient(body) {
      try {
        await fetch('/teacher/profile/data-client/save', {
          method: 'POST',
          body: JSON.stringify(body)
        });
      } catch (err) {
        console.error(err);
      }
    }
  }
});
