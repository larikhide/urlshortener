package mem

import (
	"context"
	"database/sql"
	"sync"

	"github.com/larikhide/urlshortener/app/repos/urls"
)

var _ urls.URLStore = &MemDB{}

type MemDB struct {
	sync.Mutex
	//TODO: по хорошему переписать на uuid
	//m map[uuid.UUID]urls.URL
	m map[string]urls.URL
}

func NewDB() *MemDB {
	return &MemDB{
		//m: make(map[uuid.UUID]urls.URL),
		m: make(map[string]urls.URL),
	}
}

func (m *MemDB) RetrieveInitialUrl(ctx context.Context, shortUrl string) (*urls.URL, error) {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	u, ok := m.m[shortUrl]
	if ok {
		return &u, nil
	}
	return nil, sql.ErrNoRows
}

func (m *MemDB) SaveUrlMapping(ctx context.Context, shortUrl string, longUrl string) error {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// uid := uuid.New()
	// u.ID = uid
	// m.m[u.ID] = u
	m.m[shortUrl] = urls.URL{
		ShortURL: shortUrl,
		LongURL:  longUrl,
	}

	return nil
}

// func (us *Users) Create(ctx context.Context, u user.User) (*uuid.UUID, error) {
// 	us.Lock()
// 	defer us.Unlock()

// 	select {
// 	case <-ctx.Done():
// 		return nil, ctx.Err()
// 	default:
// 	}

// 	uid := uuid.New()
// 	u.ID = uid
// 	us.m[u.ID] = u
// 	return &uid, nil
// }

// func (us *Users) Read(ctx context.Context, uid uuid.UUID) (*user.User, error) {
// 	us.Lock()
// 	defer us.Unlock()

// 	select {
// 	case <-ctx.Done():
// 		return nil, ctx.Err()
// 	default:
// 	}
// 	u, ok := us.m[uid]
// 	if ok {
// 		return &u, nil
// 	}
// 	return nil, sql.ErrNoRows
// }

// // не возвращает ошибку если не нашли
// func (us *Users) Delete(ctx context.Context, uid uuid.UUID) error {
// 	us.Lock()
// 	defer us.Unlock()

// 	select {
// 	case <-ctx.Done():
// 		return ctx.Err()
// 	default:
// 	}

// 	delete(us.m, uid)
// 	return nil
// }

// func (us *Users) SearchUsers(ctx context.Context, s string) (chan user.User, error) {
// 	us.Lock()
// 	defer us.Unlock()

// 	select {
// 	case <-ctx.Done():
// 		return nil, ctx.Err()
// 	default:
// 	}

// 	chout := make(chan user.User, 100)

// 	go func() {
// 		defer close(chout)
// 		us.Lock()
// 		defer us.Unlock()
// 		for _, u := range us.m {
// 			if strings.Contains(u.Name, s) {
// 				select {
// 				case <-ctx.Done():
// 					return
// 				case <-time.After(2 * time.Second):
// 					return
// 				case chout <- u:
// 				}
// 			}
// 		}
// 	}()

// 	return chout, nil
// }
