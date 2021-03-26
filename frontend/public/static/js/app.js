var ws;

((async () => {
  const events = {
    readingList: {
      saved: function (ns, msg) {
        app.books = msg.unmarshal()
        // or make a new http fetch
        // fetchTodos(function (items) {
        //   app.todos = msg.unmarshal()
        // });
      }
    }
  };

  const conn = await neffos.dial("ws://127.0.0.1:8080/readinglist/sync", events);
  ws = await conn.connect("readinglist");
})()).catch(console.error);



function fetchbooks(onComplete) {
  axios.get("/api/readinglist").then(response => {
    if (response.data === null) {
      return;
    }
    console.log("we are here")
    onComplete(response.data.data);
    console.log('adfdsfds')
  });
}

var bookStorage = {
  fetch: function () {
    var books = [];
    console.log(books)
    fetchbooks(function (items) {
      for (var i = 0; i < items.length; i++) {
        console.log(items[i])
        books.push(items[i]);
      }
    });
    console.log(books)
    return books;
  },
  save: function (books) {
    console.log(JSON.stringify(books))
    console.log(books)
    console.log(JSON.stringify(books[0]))
    console.log(books[0])
    req = {
        "test" : "test"
      }
    console.log(req)
    axios.post("/api/readinglist", req).then(response => {
      if (!response.data.success) {
        window.alert("saving had a failure");
        return;
      }
      console.log("send: save");
      ws.Emit("save")
    });
  }
}

// filters
var filters = {
  all: function (books) {
    return books
  },
  active: function (books) {
    return books.filter(function (book) {
      return !book.completed
    })
  },
  completed: function (books) {
    return books.filter(function (book) {
      return book.completed
    })
  }
}

// app Vue instance
var app = new Vue({
  // app initial state
  data: {
    books: bookStorage.fetch(),
    newBook: '',
    editedBook: null,
    visibility: 'all'
  },

  computed: {
    filteredBooks: function () {
      return filters[this.visibility](this.books)
    },
    remaining: function () {
      return filters.active(this.books).length
    },
    allDone: {
      get: function () {
        return this.remaining === 0
      },
      set: function (value) {
        this.books.forEach(function (book) {
          book.completed = value
        })
        this.notifyChange();
      }
    }
  },

  filters: {
    pluralize: function (n) {
      return n === 1 ? 'book' : 'books'
    }
  },


  methods: {
    notifyChange: function () {
      bookStorage.save(this.books)
    },
    addBook: function () {
      var value = this.newBook && this.newBook.trim()
      if (!value) {
        return
      }
      this.books.push({
        id: this.books.length + 1, // just for the client-side.
        title: value,
        completed: false
      })
      this.newBook = ''
      this.notifyChange();
    },

    completeBook: function (book) {
      if (book.completed) {
        book.completed = false;
      } else {
        book.completed = true;
      }
      this.notifyChange();
    },
    removeBook: function (book) {
      this.books.splice(this.books.indexOf(book), 1)
      this.notifyChange();
    },

    editBook: function (book) {
      this.beforeEditTitle = book.title
      this.beforeEditAuthor = book.author
      this.editedBook = book
    },

    doneEdit: function (book) {
      if (!this.editedBook) {
        return
      }
      this.editedBook = null
      book.title = book.title.trim();
      if (!book.title) {
        this.removeBook(book);
      }
      this.notifyChange();
    },

    cancelEdit: function (book) {
      this.editedBook = null
      book.title = this.beforeEditTitle
      book.author = this.beforeEditAuthor
    },

    removeCompleted: function () {
      this.books = filters.active(this.books);
      this.notifyChange();
    }
  },

  directives: {
    'book-focus': function (el, binding) {
      if (binding.value) {
        el.focus()
      }
    }
  }
})

// handle routing
function onHashChange() {
  var visibility = window.location.hash.replace(/#\/?/, '')
  if (filters[visibility]) {
    app.visibility = visibility
  } else {
    window.location.hash = ''
    app.visibility = 'all'
  }
}

window.addEventListener('hashchange', onHashChange)
onHashChange()

// mount
app.$mount('.readinglistapp')