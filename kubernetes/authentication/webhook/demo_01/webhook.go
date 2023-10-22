package main

import (
	"encoding/json"
	"github.com/golang/glog"
	authentication "k8s.io/api/authentication/v1beta1"
	"k8s.io/klog/v2"
	"net/http"
	"strings"
)

type WebHookServer struct {
	server *http.Server
}

func (ctx *WebHookServer) serve(w http.ResponseWriter, r *http.Request) {
	// 反序列化APIServer发送的TokenReview对象，注意APIServer可能是发送authentication.k8s.io/v1，也可能是authentication.k8s.io/v1beta1版本
	var req authentication.TokenReview
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		// 如果反序列化失败，那么认为认证失败
		klog.Error(err, "decoder request body error.")
		req.Status = authentication.TokenReviewStatus{Authenticated: false}
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(req)
		return
	}
	// token的格式必须是webhook服务可以正确解析的
	// 判断token是否包含':' ,如果不包含，则返回认证失败
	if !(strings.Contains(req.Spec.Token, ":")) {
		klog.Error(err, "token invalied.")
		req.Status = authentication.TokenReviewStatus{Authenticated: false}
		//req.Status = map[string]interface{}{"authenticated": false}
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(req)
		return
	}
	// split token, 获取type
	tokenSlice := strings.SplitN(req.Spec.Token, ":", -1)
	glog.Infof("tokenSlice: ", tokenSlice)
	hookType := tokenSlice[0]
	switch hookType {
	case "github":
		githubToken := tokenSlice[1]
		err := authByGithub(githubToken)
		if err != nil {
			klog.Error(err, "auth by github error")
			req.Status = authentication.TokenReviewStatus{Authenticated: false}
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(req)
			return
		}
		klog.Info("auth by github success")
		req.Status = authentication.TokenReviewStatus{Authenticated: true}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(req)
		return
	case "ldap":
		username := tokenSlice[1]
		password := tokenSlice[2]
		err := authByLdap(username, password)
		if err != nil {
			klog.Error(err, "auth by ldap error")
			req.Status = authentication.TokenReviewStatus{Authenticated: false}
			//req.Status = map[string]interface{}{"authenticated": false}
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(req)
			return
		}
		klog.Info("auth by ldap success")
		req.Status = authentication.TokenReviewStatus{Authenticated: true}
		//req.Status = map[string]interface{}{"authenticated": true}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(req)
		return
	}
}
