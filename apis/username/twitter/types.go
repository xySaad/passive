package twitter

type Resp struct {
	Data struct {
		User struct {
			Result struct {
				Typename string `json:"__typename"`
				Core     struct {
					Name string `json:"name"`
				} `json:"core"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}
