package com.github.cheejyg.conv.java;

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import com.github.cheejyg.conv.java.*;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * @author Cheejyg
 */
public class Conv {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ToString
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ToString parses <code>Object</code> object_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param object_ The <code>Object</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static String ToString(Object object_) {
		return object_.toString();
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Array conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * BooleansToBooleanArray parses <code>Boolean[]</code> booleans_ to
	 * <code>Boolean[]</code>.
	 * 
	 * @author Cheejyg
	 * @param booleans_ The <code>Boolean[]</code> to be parsed.
	 * @return <code>Boolean[]</code>
	 */
	public static Boolean[] BooleansToBooleanArray(Boolean... booleans_) {
		return booleans_;
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Int conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * IntToString parses <code>int</code> int_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static String IntToString(int int_) {
		String string_ = ToString(int_);

		return string_;
	}

	/**
	 * IntToLong parses <code>int</code> int_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static Long IntToLong(int int_) {
		Long long_ = new Long(int_);

		return long_;
	}

	/**
	 * IntToInteger parses <code>int</code> int_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static Integer IntToInteger(int int_) {
		return int_;
	}

	/**
	 * IntToShort parses <code>int</code> int_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static Short IntToShort(int int_) {
		Short short_ = new Short((short) int_);

		return short_;
	}

	/**
	 * IntToByte parses <code>int</code> int_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static Byte IntToByte(int int_) {
		Byte byte_ = new Byte((byte) int_);

		return byte_;
	}

	/**
	 * IntToBoolean parses <code>int</code> int_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static Boolean IntToBoolean(int int_) {
		if (int_ == 0) {
			return false;
		} else if (int_ == 1) {
			return true;
		} else {
		}

		return null;
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Short conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ShortToString parses <code>short</code> short_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static String ShortToString(short short_) {
		String string_ = ToString(short_);

		return string_;
	}

	/**
	 * ShortToLong parses <code>short</code> short_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static Long ShortToLong(short short_) {
		Long long_ = new Long(short_);

		return long_;
	}

	/**
	 * ShortToInteger parses <code>short</code> short_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static Integer ShortToInteger(short short_) {
		Integer integer_ = new Integer(short_);

		return integer_;
	}

	/**
	 * ShortToShort parses <code>short</code> short_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static Short ShortToShort(short short_) {
		return short_;
	}

	/**
	 * ShortToByte parses <code>short</code> short_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static Byte ShortToByte(short short_) {
		Byte byte_ = new Byte((byte) short_);

		return byte_;
	}

	/**
	 * ShortToBoolean parses <code>short</code> short_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static Boolean ShortToBoolean(short short_) {
		if (short_ == 0) {
			return false;
		} else if (short_ == 1) {
			return true;
		} else {
		}

		return null;
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Byte conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ByteToString parses <code>byte</code> byte_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static String ByteToString(byte byte_) {
		String string_ = ToString(byte_);

		return string_;
	}

	/**
	 * ByteToLong parses <code>byte</code> byte_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static Long ByteToLong(byte byte_) {
		Long long_ = new Long(byte_);

		return long_;
	}

	/**
	 * ByteToInteger parses <code>byte</code> byte_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static Integer ByteToInteger(byte byte_) {
		Integer integer_ = new Integer(byte_);

		return integer_;
	}

	/**
	 * ByteToShort parses <code>byte</code> byte_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static Short ByteToShort(byte byte_) {
		Short short_ = new Short(byte_);

		return short_;
	}

	/**
	 * ByteToByte returns <code>Byte</code> byte_.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static Byte ByteToByte(byte byte_) {
		return byte_;
	}

	/**
	 * ByteToBoolean parses <code>byte</code> byte_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static Boolean ByteToBoolean(byte byte_) {
		if (byte_ == 0) {
			return false;
		} else if (byte_ == 1) {
			return true;
		} else {
		}

		return null;
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Boolean conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * BooleanToString parses <code>boolean</code> boolean_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static String BooleanToString(boolean boolean_) {
		if (boolean_ == false) {
			return "false";
		} else if (boolean_ == true) {
			return "true";
		} else {
		}

		return null;
	}

	/**
	 * BooleanToLong parses <code>boolean</code> boolean_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static Long BooleanToLong(boolean boolean_) {
		if (boolean_ == false) {
			return 0L;
		} else if (boolean_ == true) {
			return 1L;
		} else {
		}

		return null;
	}

	/**
	 * BooleanToInteger parses <code>boolean</code> boolean_ to
	 * <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static Integer BooleanToInteger(boolean boolean_) {
		if (boolean_ == false) {
			return 0;
		} else if (boolean_ == true) {
			return 1;
		} else {
		}

		return null;
	}

	/**
	 * BooleanToShort parses <code>boolean</code> boolean_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static Short BooleanToShort(boolean boolean_) {
		if (boolean_ == false) {
			return 0;
		} else if (boolean_ == true) {
			return 1;
		} else {
		}

		return null;
	}

	/**
	 * BooleanToByte parses <code>boolean</code> boolean_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static Byte BooleanToByte(boolean boolean_) {
		if (boolean_ == false) {
			return 0;
		} else if (boolean_ == true) {
			return 1;
		} else {
		}

		return null;
	}

	/**
	 * BooleanToBoolean returns <code>Boolean</code> boolean_.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be returned.
	 * @return <code>Boolean</code>
	 */
	public static Boolean BooleanToBoolean(boolean boolean_) {
		return boolean_;
	}
}
