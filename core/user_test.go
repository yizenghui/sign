package core

import (
	"testing"

	wxbizdatacrypt "github.com/yilee/wx-biz-data-crypt"
)

func Test_GetOpenID(t *testing.T) {
	u2, _ := GetOpenID("00397Gwx0b39sj1Pm3wx0guNwx097GwP")
	t.Fatal(u2)
}

func Test_SendPostUpdateMSG(t *testing.T) {
	u2 := SendPostUpdateMSG("oe9Hq0GwS4umXENTCn4lJgxaNVog", "af1a0cdf6adfbf4030358fc2b4264d24", "tttt", "")
	t.Fatal(u2)
}

func Test_GetwxCodeUnlimit(t *testing.T) {
	u2, err := GetwxCodeUnlimit("123", "")
	t.Fatal(u2)
	t.Fatal(err)
}

func Test_GetToken(t *testing.T) {

	token, _ := TokenServe.Token()
	t.Fatal(token)
}

func Test_GetCryptData(t *testing.T) {
	sessionKey := `YGTTNkGX+yJZuaDklZjhJQ==`
	iv := `7hrhk+Zci2YCuUOmu73UWg==`
	encryptedData := `CbMkTL1kJXhSaXXlr367ckPfryMWNA8lquVJL9MnFr9kzI6XCLSUm2h0l+VXQp+FCGMe2muuLDy9q6C89gekTDsm200GYvRD+Br5eajKFrrUJnGNSWxV1VtsdMxW0AeVMyrRBTPKvGVabyT1jM0fZ4SWNgu5sdYgO2HrASKN7fRS2QJtkGlHChcsketyP3sfph/Jsm0QmZl3xEVPfZJSnDWl61mPemcYS/eaULFlKx5aeIdaPgX+6yFX1lBzbHj6Sfr+tx00GXHS9RnDObObT1+R5A6yRUDUmVZfXiGGID1E2aSY85EuE9L9wdlVrUzDMTLS9AVyRMjhHgA34EV+rMXQPgggRHsDqdF5rhmxKsHWJZzR0aZPYriCM+m1YLY0cZUsBPe5Dx+KcJuThSC2zj+/CUW8rtkwVywnMouo3pchzNbNi4EN/zdYypHEXV5HNarix4TG1jZadoJh5/6MjwuWDz4rzmtl4VJ244gdbGE=`
	pc := wxbizdatacrypt.NewWXBizDataCrypt(`wxa94ddd94358b2d1d`, sessionKey)
	userInfo, err := pc.Decrypt(encryptedData, iv)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(userInfo)
}
