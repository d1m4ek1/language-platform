import { defineStore } from 'pinia';

type State = {
  dataClient: IClient
  isGetDataClient: boolean
}

type Getters = {
  getFullName(state: State): string
  getAvatar(state: State): string
  getTextIsTeacher(state: State): string
  getAboutMe(state: State): string
  getLanguageItems(state: State): string[]
  getDate(state: State): string
}

type Actions = {
  sendGetDataClient(): Promise<IStatusResponse>
}

export const useClient = defineStore<string, State, Getters, Actions>('client', {
  state: ():State => ({
    dataClient: {
      firstName: "",
      lastName: "",
      aboutMe: "",
      regDate: "",
      avatar: "",
      isTeacher: null,
      interfaceLang: "ru",
      languageItems: [],
    },
    isGetDataClient: false,
  }),
  getters: {
    getFullName: (state): string => `${state.dataClient.firstName} ${state.dataClient.lastName}`,
    getAvatar: (state): string => state.dataClient.avatar || "/default/avatar/1.png",
    getTextIsTeacher: (state): string => state.dataClient.isTeacher ? "Преподаватель" : "Ученик",
    getAboutMe: (state): string => state.dataClient.aboutMe || "Вы о себе ни чего не написали",
    getLanguageItems: (state): string[] => {
      const langItems: ILanguageItems[] | null = state.dataClient.languageItems
      if (langItems === null) return []

      return langItems.map(item => {
        let langWord: string = "Английский"

        switch (item.lang) {
          case "en":
            langWord = "Английский"
            break
        }

        return `${langWord} - ${item.lvl}`
      })
    },
    getDate: (state): string => {
      let date: Date = new Date(state.dataClient.regDate);
      return date.toLocaleDateString('ru-RU', {
        weekday: "long",
        year: 'numeric',
        month: 'long',
        day: 'numeric'
      });
    },
  },
  actions: {
    async sendGetDataClient(): Promise<IStatusResponse> {
      try {
        const response = await fetch('/student/profile/data-client')
        const json: IGetDataClient = await response.json()

        const dataClient: IClient = json.dataClient

        Object.keys(this.dataClient).forEach((key: string) => {
          if (this.dataClient[key] === undefined) return
          this.dataClient[key] = dataClient[key];
        });

        this.isGetDataClient = true;

        return json.responseStatus
      } catch (e) {
        console.error(e)

        return {
          success: false,
          message: "Неизвестная ошибка"
        }
      }
    }
  }
})