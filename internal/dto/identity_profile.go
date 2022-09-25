package dto

import (
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/firebase_auth"
)

func ToAppIdentityProfile(ur firebase_auth.FirebaseProfile) entity.IdentityProfile {
	claims := map[string]bool{}

	for k, v := range ur.CustomClaims {
		if value, ok := v.(bool); ok {
			claims[k] = value
		}
	}

	return entity.IdentityProfile{
		UID:    ur.UID,
		Email:  ur.Email,
		Claims: claims,
	}
}
